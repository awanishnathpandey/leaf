-- name: GetEmailTemplateByName :one
SELECT name, content
FROM email_templates
WHERE name = $1
LIMIT 1;