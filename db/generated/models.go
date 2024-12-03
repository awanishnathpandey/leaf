// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package generated

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type EmailTemplate struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Content   string   `json:"content"`
	MailTo    []string `json:"mail_to"`
	MailCc    []string `json:"mail_cc"`
	MailBcc   []string `json:"mail_bcc"`
	MailData  []byte   `json:"mail_data"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

type File struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	FilePath     string `json:"file_path"`
	FileType     string `json:"file_type"`
	FileBytes    int64  `json:"file_bytes"`
	AutoDownload bool   `json:"auto_download"`
	FolderID     int64  `json:"folder_id"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
}

type Folder struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

type Group struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

type GroupFile struct {
	GroupID   int64  `json:"group_id"`
	FileID    int64  `json:"file_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type GroupFolder struct {
	GroupID   int64  `json:"group_id"`
	FolderID  int64  `json:"folder_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type GroupUser struct {
	GroupID   int64  `json:"group_id"`
	UserID    int64  `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type PasswordReset struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	ResetToken string `json:"reset_token"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

type Permission struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

type RolePermission struct {
	RoleID       int64  `json:"role_id"`
	PermissionID int64  `json:"permission_id"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
}

type User struct {
	ID              int64       `json:"id"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	Email           string      `json:"email"`
	Password        string      `json:"password"`
	JobTitle        pgtype.Text `json:"job_title"`
	LineOfBusiness  pgtype.Text `json:"line_of_business"`
	LineManager     pgtype.Text `json:"line_manager"`
	EmailVerifiedAt pgtype.Int8 `json:"email_verified_at"`
	LastSeenAt      int64       `json:"last_seen_at"`
	CreatedAt       int64       `json:"created_at"`
	UpdatedAt       int64       `json:"updated_at"`
	DeletedAt       pgtype.Int8 `json:"deleted_at"`
	CreatedBy       string      `json:"created_by"`
	UpdatedBy       string      `json:"updated_by"`
}

type UserRole struct {
	UserID    int64  `json:"user_id"`
	RoleID    int64  `json:"role_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
