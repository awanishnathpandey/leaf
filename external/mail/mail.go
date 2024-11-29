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
func (ms *MailService) SendEmail(to, subject, templateName string, data map[string]interface{}) error {
	// Render the template with dynamic data
	renderedTemplate, err := RenderTemplate(templateName, data)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	// Create a new email message
	msg := mail.NewMsg()

	// Set email metadata
	if err := msg.From(ms.config.SMTPFrom); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	if err := msg.To(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	// Set email content (HTML format)
	msg.Subject(subject)
	msg.SetBodyString(mail.TypeTextHTML, renderedTemplate)

	// Send the email
	ctx, cancel := context.WithTimeout(context.Background(), ms.config.SendTimeout)
	defer cancel()
	// fmt.Println("SMTP User:", ms.config.SMTPUser)
	// fmt.Println("SMTP From:", ms.config.SMTPFrom)
	// fmt.Println("SMTP Host:", ms.config.SMTPHost)

	if err := ms.client.DialAndSendWithContext(ctx, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
