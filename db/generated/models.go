// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package generated

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AppConfig struct {
	ID         int64  `db:"id" json:"id"`
	ConfigKey  string `db:"config_key" json:"config_key"`
	ConfigData []byte `db:"config_data" json:"config_data"`
	CreatedAt  int64  `db:"created_at" json:"created_at"`
	UpdatedAt  int64  `db:"updated_at" json:"updated_at"`
	CreatedBy  string `db:"created_by" json:"created_by"`
	UpdatedBy  string `db:"updated_by" json:"updated_by"`
}

type AuditLog struct {
	ID          int64  `db:"id" json:"id"`
	TableName   string `db:"table_name" json:"table_name"`
	Actor       string `db:"actor" json:"actor"`
	Action      string `db:"action" json:"action"`
	IpAddress   string `db:"ip_address" json:"ip_address"`
	RecordKey   string `db:"record_key" json:"record_key"`
	Description string `db:"description" json:"description"`
	Timestamp   int64  `db:"timestamp" json:"timestamp"`
}

type CronJob struct {
	ID          int64       `db:"id" json:"id"`
	Slug        string      `db:"slug" json:"slug"`
	Name        string      `db:"name" json:"name"`
	Schedule    string      `db:"schedule" json:"schedule"`
	Active      pgtype.Bool `db:"active" json:"active"`
	Description string      `db:"description" json:"description"`
	LastRunAt   int64       `db:"last_run_at" json:"last_run_at"`
	CreatedAt   int64       `db:"created_at" json:"created_at"`
	UpdatedAt   int64       `db:"updated_at" json:"updated_at"`
	CreatedBy   string      `db:"created_by" json:"created_by"`
	UpdatedBy   string      `db:"updated_by" json:"updated_by"`
}

type CronJobLog struct {
	ID              int64  `db:"id" json:"id"`
	CronSlug        string `db:"cron_slug" json:"cron_slug"`
	Status          string `db:"status" json:"status"`
	Message         string `db:"message" json:"message"`
	StartTime       int64  `db:"start_time" json:"start_time"`
	EndTime         int64  `db:"end_time" json:"end_time"`
	AffectedRecords int64  `db:"affected_records" json:"affected_records"`
}

type EmailTemplate struct {
	ID        int64    `db:"id" json:"id"`
	Name      string   `db:"name" json:"name"`
	Content   string   `db:"content" json:"content"`
	MailTo    []string `db:"mail_to" json:"mail_to"`
	MailCc    []string `db:"mail_cc" json:"mail_cc"`
	MailBcc   []string `db:"mail_bcc" json:"mail_bcc"`
	MailData  []byte   `db:"mail_data" json:"mail_data"`
	CreatedAt int64    `db:"created_at" json:"created_at"`
	UpdatedAt int64    `db:"updated_at" json:"updated_at"`
}

type File struct {
	ID              int64  `db:"id" json:"id"`
	Name            string `db:"name" json:"name"`
	Slug            string `db:"slug" json:"slug"`
	FilePath        string `db:"file_path" json:"file_path"`
	FileType        string `db:"file_type" json:"file_type"`
	FileBytes       int64  `db:"file_bytes" json:"file_bytes"`
	FileContentType string `db:"file_content_type" json:"file_content_type"`
	AutoDownload    bool   `db:"auto_download" json:"auto_download"`
	FolderID        int64  `db:"folder_id" json:"folder_id"`
	CreatedAt       int64  `db:"created_at" json:"created_at"`
	UpdatedAt       int64  `db:"updated_at" json:"updated_at"`
	CreatedBy       string `db:"created_by" json:"created_by"`
	UpdatedBy       string `db:"updated_by" json:"updated_by"`
}

type Folder struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Slug        string `db:"slug" json:"slug"`
	Description string `db:"description" json:"description"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedBy   string `db:"created_by" json:"created_by"`
	UpdatedBy   string `db:"updated_by" json:"updated_by"`
}

type Group struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedBy   string `db:"created_by" json:"created_by"`
	UpdatedBy   string `db:"updated_by" json:"updated_by"`
}

type GroupFile struct {
	GroupID   int64  `db:"group_id" json:"group_id"`
	FileID    int64  `db:"file_id" json:"file_id"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"created_by"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}

type GroupFolder struct {
	GroupID   int64  `db:"group_id" json:"group_id"`
	FolderID  int64  `db:"folder_id" json:"folder_id"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"created_by"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}

type GroupUser struct {
	GroupID   int64  `db:"group_id" json:"group_id"`
	UserID    int64  `db:"user_id" json:"user_id"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"created_by"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}

type Notification struct {
	ID                 int64   `db:"id" json:"id"`
	NotificationType   string  `db:"notification_type" json:"notification_type"`
	RecordKeyID        int64   `db:"record_key_id" json:"record_key_id"`
	Payload            []byte  `db:"payload" json:"payload"`
	StartTimeAt        int64   `db:"start_time_at" json:"start_time_at"`
	EndTimeAt          int64   `db:"end_time_at" json:"end_time_at"`
	IsPushNotification bool    `db:"is_push_notification" json:"is_push_notification"`
	Status             string  `db:"status" json:"status"`
	GroupIds           []int64 `db:"group_ids" json:"group_ids"`
	UserIds            []int64 `db:"user_ids" json:"user_ids"`
	CreatedAt          int64   `db:"created_at" json:"created_at"`
	CreatedBy          string  `db:"created_by" json:"created_by"`
}

type NotificationTemplate struct {
	ID              int64    `db:"id" json:"id"`
	Title           string   `db:"title" json:"title"`
	Body            string   `db:"body" json:"body"`
	Description     string   `db:"description" json:"description"`
	ResponseOptions []string `db:"response_options" json:"response_options"`
	CreatedAt       int64    `db:"created_at" json:"created_at"`
	UpdatedAt       int64    `db:"updated_at" json:"updated_at"`
	CreatedBy       string   `db:"created_by" json:"created_by"`
	UpdatedBy       string   `db:"updated_by" json:"updated_by"`
}

type PasswordReset struct {
	ID         int64  `db:"id" json:"id"`
	UserID     int64  `db:"user_id" json:"user_id"`
	ResetToken string `db:"reset_token" json:"reset_token"`
	CreatedAt  int64  `db:"created_at" json:"created_at"`
	UpdatedAt  int64  `db:"updated_at" json:"updated_at"`
	CreatedBy  string `db:"created_by" json:"created_by"`
	UpdatedBy  string `db:"updated_by" json:"updated_by"`
}

type Permission struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedBy   string `db:"created_by" json:"created_by"`
	UpdatedBy   string `db:"updated_by" json:"updated_by"`
}

type Role struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedBy   string `db:"created_by" json:"created_by"`
	UpdatedBy   string `db:"updated_by" json:"updated_by"`
}

type RolePermission struct {
	RoleID       int64  `db:"role_id" json:"role_id"`
	PermissionID int64  `db:"permission_id" json:"permission_id"`
	CreatedAt    int64  `db:"created_at" json:"created_at"`
	UpdatedAt    int64  `db:"updated_at" json:"updated_at"`
	CreatedBy    string `db:"created_by" json:"created_by"`
	UpdatedBy    string `db:"updated_by" json:"updated_by"`
}

type User struct {
	ID                     int64       `db:"id" json:"id"`
	FirstName              string      `db:"first_name" json:"first_name"`
	LastName               string      `db:"last_name" json:"last_name"`
	Email                  string      `db:"email" json:"email"`
	Password               string      `db:"password" json:"password"`
	JobTitle               pgtype.Text `db:"job_title" json:"job_title"`
	LineOfBusiness         pgtype.Text `db:"line_of_business" json:"line_of_business"`
	LineManager            pgtype.Text `db:"line_manager" json:"line_manager"`
	EmailVerifiedAt        pgtype.Int8 `db:"email_verified_at" json:"email_verified_at"`
	LastSeenAt             int64       `db:"last_seen_at" json:"last_seen_at"`
	LastNotificationReadAt int64       `db:"last_notification_read_at" json:"last_notification_read_at"`
	CreatedAt              int64       `db:"created_at" json:"created_at"`
	UpdatedAt              int64       `db:"updated_at" json:"updated_at"`
	DeletedAt              pgtype.Int8 `db:"deleted_at" json:"deleted_at"`
	CreatedBy              string      `db:"created_by" json:"created_by"`
	UpdatedBy              string      `db:"updated_by" json:"updated_by"`
}

type UserNotificationResponse struct {
	ID             int64  `db:"id" json:"id"`
	NotificationID int64  `db:"notification_id" json:"notification_id"`
	UserID         int64  `db:"user_id" json:"user_id"`
	Response       string `db:"response" json:"response"`
	CreatedAt      int64  `db:"created_at" json:"created_at"`
	CreatedBy      string `db:"created_by" json:"created_by"`
}

type UserRole struct {
	UserID    int64  `db:"user_id" json:"user_id"`
	RoleID    int64  `db:"role_id" json:"role_id"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"created_by"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}
