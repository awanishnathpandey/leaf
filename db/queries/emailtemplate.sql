-- name: GetEmailTemplateByName :one
SELECT *
FROM email_templates
WHERE name = $1
LIMIT 1;