// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package generated

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Folder struct {
	ID          int32              `json:"id"`
	Name        string             `json:"name"`
	Slug        string             `json:"slug"`
	Description string             `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	CreatedBy   string             `json:"created_by"`
	UpdatedBy   string             `json:"updated_by"`
}

type User struct {
	ID              int32              `json:"id"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Email           string             `json:"email"`
	Password        string             `json:"password"`
	JobTitle        pgtype.Text        `json:"job_title"`
	LineOfBusiness  pgtype.Text        `json:"line_of_business"`
	EmailVerifiedAt pgtype.Timestamptz `json:"email_verified_at"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
	DeletedAt       pgtype.Timestamptz `json:"deleted_at"`
	CreatedBy       string             `json:"created_by"`
	UpdatedBy       string             `json:"updated_by"`
	DeletedBy       pgtype.Text        `json:"deleted_by"`
}
