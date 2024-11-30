package mail

import (
	"context"
	"fmt"
	"strings"

	mail "github.com/wneessen/go-mail"
)

// MailService encapsulates the email configuration and client
type MailService struct {
	config Config
	client *mail.Client
}

// NewMailService initializes the email service with configuration
func NewMailService(cfg Config) (*MailService, error) {
	client, err := mail.NewClient(
		cfg.SMTPHost,
		mail.WithPort(cfg.SMTPPort),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.SMTPUser),
		mail.WithPassword(cfg.SMTPPass),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithTimeout(cfg.ConnectTimeout),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create SMTP client: %w", err)
	}

	return &MailService{config: cfg, client: client}, nil
}

// ReplacePlaceholders replaces placeholders in the email body with dynamic values
func ReplacePlaceholders(body string, data map[string]string) string {
	for placeholder, value := range data {
		placeholderFormat := fmt.Sprintf("{{%s}}", placeholder)
		body = strings.ReplaceAll(body, placeholderFormat, value)
	}
	return body
}

// SendEmail sends an email using the template and dynamic data
func (ms *MailService) SendEmail(to []string, cc []string, bcc []string, subject, templateContent string, data map[string]interface{}) error {
	// Render the template with dynamic data
	renderedTemplate, err := RenderTemplate(templateContent, data)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	// Create a new email message
	msg := mail.NewMsg()

	// Set email metadata
	if err := msg.From(ms.config.SMTPFrom); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	var errorRecipients []string
	// // Add multiple recipients
	for _, recipient := range to {
		if err := msg.AddTo(recipient); err != nil {
			// Log the error but continue with the next recipient
			errorRecipients = append(errorRecipients, fmt.Sprintf("failed to add To recipient %s: %v", recipient, err))
		}
	}

	for _, recipient := range cc {
		if err := msg.AddCc(recipient); err != nil {
			// Log the error but continue with the next recipient
			errorRecipients = append(errorRecipients, fmt.Sprintf("failed to add Cc recipient %s: %v", recipient, err))
		}
	}

	for _, recipient := range bcc {
		if err := msg.AddBcc(recipient); err != nil {
			// Log the error but continue with the next recipient
			errorRecipients = append(errorRecipients, fmt.Sprintf("failed to add Bcc recipient %s: %v", recipient, err))
		}
	}
	// if err := msg.To("res@res.com,comma@res.com"); err != nil {
	// 	return fmt.Errorf("failed to set recipient: %w", err)
	// }
	// If there were any errors, log them but still proceed to send the email
	if len(errorRecipients) > 0 {
		for _, msg := range errorRecipients {
			// Optionally, log each error
			fmt.Println(msg)
		}
	}

	// Set email content (HTML format)
	msg.Subject(subject)
	msg.SetBodyString(mail.TypeTextHTML, renderedTemplate)

	// Send the email
	ctx, cancel := context.WithTimeout(context.Background(), ms.config.SendTimeout)
	defer cancel()
	// fmt.Println("SMTP User:", ms.config.SMTPUser)
	// fmt.Println("SMTP From:", ms.config.SMTPFrom)
	// fmt.Println("SMTP to:", to)

	if err := ms.client.DialAndSendWithContext(ctx, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
