package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

func RenderTemplate(templateContent string, data map[string]interface{}) (string, error) {
	// Parse the template content (HTML with placeholders)
	tmpl, err := template.New("email").Parse(templateContent)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute the template with the provided data
	var renderedTemplate bytes.Buffer
	err = tmpl.Execute(&renderedTemplate, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	// Return the rendered template content
	return renderedTemplate.String(), nil
}

func ConvertMailData(mailData []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(mailData, &result) // Unmarshal JSON data into a map
	if err != nil {
		return nil, err
	}
	return result, nil
}
