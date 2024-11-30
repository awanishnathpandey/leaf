package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"fmt"

	"github.com/awanishnathpandey/leaf/external/mail"
	"github.com/awanishnathpandey/leaf/graph/model"
	"github.com/awanishnathpandey/leaf/internal/utils"
)

// SendEmail is the resolver for the sendEmail field.
func (r *mutationResolver) SendEmail(ctx context.Context, input model.SendEmailInput) (*model.EmailResponse, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Step 1: Validate template name
	templateName := input.TemplateName
	if templateName == "" {
		return nil, fmt.Errorf("template name cannot be empty")
	}
	// Fetch the template from the database
	templateData, err := r.DB.GetEmailTemplateByName(ctx, templateName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch template: %w", err)
	}

	mailData, err := mail.ConvertMailData(templateData.MailData)
	if err != nil {
		return nil, fmt.Errorf("failed to convert mail data: %w", err)
	}

	// fmt.Println(mailData["Name"])
	// fmt.Println(mailData["Email"])
	// mailData is already of type map[string]interface{}
	if mailData != nil {
		// Update the name and email fields in the map
		mailData["Name"] = ctx.Value("userEmail").(string)  // Replace with the new name
		mailData["Email"] = ctx.Value("userEmail").(string) // Replace with the new email
	} else {
		return nil, fmt.Errorf("mailData is nil")
	}

	// Step 3: Render the template with provided data
	renderedTemplate, err := mail.RenderTemplate(templateData.Content, mailData)
	if err != nil {
		return nil, fmt.Errorf("failed to render template: %w", err)
	}

	mailSubject := "Welcome!"

	// Step 4: Send the email with the rendered template
	err = r.MailService.SendEmail(templateData.MailTo, templateData.MailCc, templateData.MailBcc, mailSubject, renderedTemplate, mailData)
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}

	// Step 5: Return response
	return &model.EmailResponse{
		Success: true,
		Message: "Email sent successfully",
	}, nil
}
