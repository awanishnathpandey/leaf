// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: emailtemplate.sql

package generated

import (
	"context"
)

const getEmailTemplateByName = `-- name: GetEmailTemplateByName :one
SELECT id, name, content, mail_to, mail_cc, mail_bcc, mail_data, created_at, updated_at
FROM email_templates
WHERE name = $1
LIMIT 1
`

func (q *Queries) GetEmailTemplateByName(ctx context.Context, name string) (EmailTemplate, error) {
	row := q.db.QueryRow(ctx, getEmailTemplateByName, name)
	var i EmailTemplate
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Content,
		&i.MailTo,
		&i.MailCc,
		&i.MailBcc,
		&i.MailData,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}