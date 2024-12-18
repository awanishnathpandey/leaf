// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AppConfig struct {
	ID         int64  `json:"id"`
	ConfigKey  string `json:"configKey"`
	ConfigData any    `json:"configData"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
	CreatedBy  string `json:"createdBy"`
	UpdatedBy  string `json:"updatedBy"`
}

type AppNotificationPayload struct {
	NotificationID  int64     `json:"notificationId"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Description     string    `json:"description"`
	ResponseOptions []*string `json:"responseOptions,omitempty"`
}

type AuditLog struct {
	ID          int64  `json:"id"`
	TableName   string `json:"tableName"`
	Actor       string `json:"actor"`
	Action      string `json:"action"`
	IPAddress   string `json:"ipAddress"`
	RecordKey   string `json:"recordKey"`
	Description string `json:"description"`
	ActorUser   *User  `json:"actorUser"`
	Timestamp   int64  `json:"timestamp"`
}

type AuditLogConnection struct {
	TotalCount int64           `json:"totalCount"`
	Edges      []*AuditLogEdge `json:"edges"`
	PageInfo   *PageInfo       `json:"pageInfo"`
}

type AuditLogEdge struct {
	Cursor string    `json:"cursor"`
	Node   *AuditLog `json:"node"`
}

type AuditLogFilter struct {
	TableName   *string `json:"tableName,omitempty"`
	Actor       *string `json:"actor,omitempty"`
	IPAddress   *string `json:"ipAddress,omitempty"`
	Action      *string `json:"action,omitempty"`
	RecordKey   *string `json:"recordKey,omitempty"`
	Description *string `json:"description,omitempty"`
}

type AuditLogSort struct {
	Field AuditLogSortField `json:"field"`
	Order SortOrder         `json:"order"`
}

type AuthUser struct {
	ID                     int64   `json:"id"`
	FirstName              string  `json:"firstName"`
	LastName               string  `json:"lastName"`
	Email                  string  `json:"email"`
	JobTitle               *string `json:"jobTitle,omitempty"`
	LineOfBusiness         *string `json:"lineOfBusiness,omitempty"`
	LineManager            *string `json:"lineManager,omitempty"`
	LastSeenAt             int64   `json:"lastSeenAt"`
	LastNotificationReadAt int64   `json:"lastNotificationReadAt"`
	CreatedAt              int64   `json:"createdAt"`
	UpdatedAt              int64   `json:"updatedAt"`
}

type CreateFile struct {
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	FilePath        string `json:"filePath"`
	FileType        string `json:"fileType"`
	FileBytes       int64  `json:"fileBytes"`
	FileContentType string `json:"fileContentType"`
	FolderID        int64  `json:"folderId"`
}

type CreateFolder struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type CreateGroup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CronJob struct {
	ID          int64                 `json:"id"`
	Active      bool                  `json:"active"`
	Name        string                `json:"name"`
	Slug        string                `json:"slug"`
	Description string                `json:"description"`
	Schedule    string                `json:"schedule"`
	LastRunAt   int64                 `json:"lastRunAt"`
	CreatedAt   int64                 `json:"createdAt"`
	UpdatedAt   int64                 `json:"updatedAt"`
	CreatedBy   string                `json:"createdBy"`
	UpdatedBy   string                `json:"updatedBy"`
	CronJobLogs *CronJobLogConnection `json:"cronJobLogs"`
}

type CronJobConnection struct {
	TotalCount int64          `json:"totalCount"`
	Edges      []*CronJobEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
}

type CronJobEdge struct {
	Cursor string   `json:"cursor"`
	Node   *CronJob `json:"node"`
}

type CronJobFilter struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Schedule    *string `json:"schedule,omitempty"`
}

type CronJobLog struct {
	ID              int64    `json:"id"`
	CronSlug        string   `json:"cronSlug"`
	Status          string   `json:"status"`
	Message         string   `json:"message"`
	StartTime       int64    `json:"startTime"`
	EndTime         int64    `json:"endTime"`
	AffectedRecords int64    `json:"affectedRecords"`
	CronJob         *CronJob `json:"cronJob"`
}

type CronJobLogConnection struct {
	TotalCount int64             `json:"totalCount"`
	Edges      []*CronJobLogEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
}

type CronJobLogEdge struct {
	Cursor string      `json:"cursor"`
	Node   *CronJobLog `json:"node"`
}

type CronJobLogFilter struct {
	Slug    *string `json:"slug,omitempty"`
	Message *string `json:"message,omitempty"`
}

type CronJobLogSort struct {
	Field CronJobLogSortField `json:"field"`
	Order SortOrder           `json:"order"`
}

type CronJobSort struct {
	Field CronJobSortField `json:"field"`
	Order SortOrder        `json:"order"`
}

type DashboardKPICount struct {
	Users       int64 `json:"users"`
	Roles       int64 `json:"roles"`
	Permissions int64 `json:"permissions"`
	Groups      int64 `json:"groups"`
	Folders     int64 `json:"folders"`
	Files       int64 `json:"files"`
}

type EmailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type File struct {
	ID              int64            `json:"id"`
	Name            string           `json:"name"`
	Slug            string           `json:"slug"`
	FilePath        string           `json:"filePath"`
	FileType        string           `json:"fileType"`
	FileBytes       int64            `json:"fileBytes"`
	FileContentType string           `json:"fileContentType"`
	AutoDownload    bool             `json:"autoDownload"`
	FolderID        int64            `json:"folderId"`
	Folder          *Folder          `json:"folder"`
	CreatedAt       int64            `json:"createdAt"`
	UpdatedAt       int64            `json:"updatedAt"`
	CreatedBy       string           `json:"createdBy"`
	UpdatedBy       string           `json:"updatedBy"`
	Groups          *GroupConnection `json:"groups"`
}

type FileConnection struct {
	TotalCount int64       `json:"totalCount"`
	Edges      []*FileEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
}

type FileEdge struct {
	Cursor string `json:"cursor"`
	Node   *File  `json:"node"`
}

type FileFilter struct {
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

type FileNotificationPayload struct {
	FileID     int64  `json:"fileId"`
	FolderID   int64  `json:"folderId"`
	FileName   string `json:"fileName"`
	FolderName string `json:"folderName"`
	EventType  string `json:"eventType"`
}

type FileSort struct {
	Field FileSortField `json:"field"`
	Order SortOrder     `json:"order"`
}

type Folder struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	CreatedAt   int64            `json:"createdAt"`
	UpdatedAt   int64            `json:"updatedAt"`
	CreatedBy   string           `json:"createdBy"`
	UpdatedBy   string           `json:"updatedBy"`
	Groups      *GroupConnection `json:"groups"`
	Files       *FileConnection  `json:"files"`
}

type FolderConnection struct {
	TotalCount int64         `json:"totalCount"`
	Edges      []*FolderEdge `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
}

type FolderEdge struct {
	Cursor string  `json:"cursor"`
	Node   *Folder `json:"node"`
}

type FolderFilter struct {
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
}

type FolderSort struct {
	Field FolderSortField `json:"field"`
	Order SortOrder       `json:"order"`
}

type Group struct {
	ID          int64             `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	CreatedAt   int64             `json:"createdAt"`
	UpdatedAt   int64             `json:"updatedAt"`
	CreatedBy   string            `json:"createdBy"`
	UpdatedBy   string            `json:"updatedBy"`
	Users       *UserConnection   `json:"users"`
	Folders     *FolderConnection `json:"folders"`
	Files       *FileConnection   `json:"files"`
}

type GroupConnection struct {
	TotalCount int64        `json:"totalCount"`
	Edges      []*GroupEdge `json:"edges"`
	PageInfo   *PageInfo    `json:"pageInfo"`
}

type GroupEdge struct {
	Cursor string `json:"cursor"`
	Node   *Group `json:"node"`
}

type GroupFilter struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type GroupSort struct {
	Field GroupSortField `json:"field"`
	Order SortOrder      `json:"order"`
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	User         *AuthUser `json:"user"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
}

type Mutation struct {
}

type MyFile struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	FilePath        string  `json:"filePath"`
	FileType        string  `json:"fileType"`
	FileBytes       int64   `json:"fileBytes"`
	FileContentType string  `json:"fileContentType"`
	AutoDownload    bool    `json:"autoDownload"`
	IsNew           bool    `json:"isNew"`
	FolderID        int64   `json:"folderId"`
	Folder          *Folder `json:"folder"`
	CreatedAt       int64   `json:"createdAt"`
	UpdatedAt       int64   `json:"updatedAt"`
}

type MyFolder struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	HasNewFile  bool      `json:"hasNewFile"`
	CreatedAt   int64     `json:"createdAt"`
	UpdatedAt   int64     `json:"updatedAt"`
	MyFiles     []*MyFile `json:"myFiles"`
}

type Notification struct {
	ID                 int64                  `json:"id"`
	NotificationType   string                 `json:"notificationType"`
	RecordKeyID        int64                  `json:"recordKeyId"`
	Payload            map[string]interface{} `json:"payload"`
	StartTimeAt        int64                  `json:"startTimeAt"`
	EndTimeAt          int64                  `json:"endTimeAt"`
	IsPushNotification bool                   `json:"isPushNotification"`
	Status             string                 `json:"status"`
	GroupIds           []int64                `json:"groupIds,omitempty"`
	UserIds            []int64                `json:"userIds,omitempty"`
	CreatedAt          int64                  `json:"createdAt"`
	CreatedBy          string                 `json:"createdBy"`
}

type NotificationInput struct {
	NotificationType   string                 `json:"notificationType"`
	RecordKeyID        int64                  `json:"recordKeyId"`
	Payload            map[string]interface{} `json:"payload"`
	StartTimeAt        int64                  `json:"startTimeAt"`
	EndTimeAt          int64                  `json:"endTimeAt"`
	Status             string                 `json:"status"`
	IsPushNotification bool                   `json:"isPushNotification"`
	GroupIds           []int64                `json:"groupIds,omitempty"`
	UserIds            []int64                `json:"userIds,omitempty"`
}

type NotificationTemplate struct {
	ID              int64    `json:"id"`
	Title           string   `json:"title"`
	Body            string   `json:"body"`
	Description     string   `json:"description"`
	ResponseOptions []string `json:"responseOptions,omitempty"`
	CreatedAt       int64    `json:"createdAt"`
	CreatedBy       string   `json:"createdBy"`
	UpdatedAt       int64    `json:"updatedAt"`
	UpdatedBy       string   `json:"updatedBy"`
}

type NotificationTemplateInput struct {
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Description     string    `json:"description"`
	ResponseOptions []*string `json:"responseOptions,omitempty"`
}

type PageInfo struct {
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
}

type Permission struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CreatedAt   int64           `json:"createdAt"`
	UpdatedAt   int64           `json:"updatedAt"`
	CreatedBy   string          `json:"createdBy"`
	UpdatedBy   string          `json:"updatedBy"`
	Roles       *RoleConnection `json:"roles"`
}

type PermissionConnection struct {
	TotalCount int64             `json:"totalCount"`
	Edges      []*PermissionEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
}

type PermissionEdge struct {
	Cursor string      `json:"cursor"`
	Node   *Permission `json:"node"`
}

type PermissionFilter struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type PermissionSort struct {
	Field PermissionSortField `json:"field"`
	Order SortOrder           `json:"order"`
}

type Query struct {
}

type Role struct {
	ID          int64                 `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	CreatedAt   int64                 `json:"createdAt"`
	UpdatedAt   int64                 `json:"updatedAt"`
	CreatedBy   string                `json:"createdBy"`
	UpdatedBy   string                `json:"updatedBy"`
	Permissions *PermissionConnection `json:"permissions"`
	Users       *UserConnection       `json:"users"`
}

type RoleConnection struct {
	TotalCount int64       `json:"totalCount"`
	Edges      []*RoleEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
}

type RoleEdge struct {
	Cursor string `json:"cursor"`
	Node   *Role  `json:"node"`
}

type RoleFilter struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type RoleSort struct {
	Field RoleSortField `json:"field"`
	Order SortOrder     `json:"order"`
}

type SendEmailInput struct {
	TemplateName string `json:"templateName"`
}

type UpdateCronJob struct {
	Active      bool   `json:"active"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Schedule    string `json:"schedule"`
}

type UpdateFile struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	FilePath string `json:"filePath"`
}

type UpdateFolder struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type UpdateGroup struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type User struct {
	ID                     int64               `json:"id"`
	FirstName              string              `json:"firstName"`
	LastName               string              `json:"lastName"`
	Email                  string              `json:"email"`
	Password               string              `json:"password"`
	JobTitle               *string             `json:"jobTitle,omitempty"`
	LineOfBusiness         *string             `json:"lineOfBusiness,omitempty"`
	LineManager            *string             `json:"lineManager,omitempty"`
	EmailVerifiedAt        *int64              `json:"emailVerifiedAt,omitempty"`
	LastSeenAt             int64               `json:"lastSeenAt"`
	LastNotificationReadAt int64               `json:"lastNotificationReadAt"`
	CreatedAt              int64               `json:"createdAt"`
	UpdatedAt              int64               `json:"updatedAt"`
	DeletedAt              *int64              `json:"deletedAt,omitempty"`
	CreatedBy              string              `json:"createdBy"`
	UpdatedBy              string              `json:"updatedBy"`
	Groups                 *GroupConnection    `json:"groups"`
	Roles                  *RoleConnection     `json:"roles"`
	AuditLogs              *AuditLogConnection `json:"auditLogs"`
}

type UserConnection struct {
	TotalCount int64       `json:"totalCount"`
	Edges      []*UserEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
}

type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

type UserFilter struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type UserNotificationResponse struct {
	ID             int64   `json:"id"`
	NotificationID int64   `json:"notificationId"`
	UserID         int64   `json:"userId"`
	Response       *string `json:"response,omitempty"`
	CreatedAt      int64   `json:"createdAt"`
	CreatedBy      string  `json:"createdBy"`
}

type UserNotificationResponseInput struct {
	UserID         int64  `json:"userId"`
	NotificationID int64  `json:"notificationId"`
	Response       string `json:"response"`
}

type UserSort struct {
	Field UserSortField `json:"field"`
	Order SortOrder     `json:"order"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type Register struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type ResetPassword struct {
	ResetToken  string `json:"resetToken"`
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

type AuditLogSortField string

const (
	AuditLogSortFieldTablename   AuditLogSortField = "TABLENAME"
	AuditLogSortFieldActor       AuditLogSortField = "ACTOR"
	AuditLogSortFieldAction      AuditLogSortField = "ACTION"
	AuditLogSortFieldIpaddress   AuditLogSortField = "IPADDRESS"
	AuditLogSortFieldRecordkey   AuditLogSortField = "RECORDKEY"
	AuditLogSortFieldDescription AuditLogSortField = "DESCRIPTION"
	AuditLogSortFieldTimestamp   AuditLogSortField = "TIMESTAMP"
)

var AllAuditLogSortField = []AuditLogSortField{
	AuditLogSortFieldTablename,
	AuditLogSortFieldActor,
	AuditLogSortFieldAction,
	AuditLogSortFieldIpaddress,
	AuditLogSortFieldRecordkey,
	AuditLogSortFieldDescription,
	AuditLogSortFieldTimestamp,
}

func (e AuditLogSortField) IsValid() bool {
	switch e {
	case AuditLogSortFieldTablename, AuditLogSortFieldActor, AuditLogSortFieldAction, AuditLogSortFieldIpaddress, AuditLogSortFieldRecordkey, AuditLogSortFieldDescription, AuditLogSortFieldTimestamp:
		return true
	}
	return false
}

func (e AuditLogSortField) String() string {
	return string(e)
}

func (e *AuditLogSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuditLogSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuditLogSortField", str)
	}
	return nil
}

func (e AuditLogSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CronJobLogSortField string

const (
	CronJobLogSortFieldSlug    CronJobLogSortField = "SLUG"
	CronJobLogSortFieldMessage CronJobLogSortField = "MESSAGE"
)

var AllCronJobLogSortField = []CronJobLogSortField{
	CronJobLogSortFieldSlug,
	CronJobLogSortFieldMessage,
}

func (e CronJobLogSortField) IsValid() bool {
	switch e {
	case CronJobLogSortFieldSlug, CronJobLogSortFieldMessage:
		return true
	}
	return false
}

func (e CronJobLogSortField) String() string {
	return string(e)
}

func (e *CronJobLogSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CronJobLogSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CronJobLogSortField", str)
	}
	return nil
}

func (e CronJobLogSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CronJobSortField string

const (
	CronJobSortFieldName        CronJobSortField = "NAME"
	CronJobSortFieldDescription CronJobSortField = "DESCRIPTION"
	CronJobSortFieldSchedule    CronJobSortField = "SCHEDULE"
)

var AllCronJobSortField = []CronJobSortField{
	CronJobSortFieldName,
	CronJobSortFieldDescription,
	CronJobSortFieldSchedule,
}

func (e CronJobSortField) IsValid() bool {
	switch e {
	case CronJobSortFieldName, CronJobSortFieldDescription, CronJobSortFieldSchedule:
		return true
	}
	return false
}

func (e CronJobSortField) String() string {
	return string(e)
}

func (e *CronJobSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CronJobSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CronJobSortField", str)
	}
	return nil
}

func (e CronJobSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FileSortField string

const (
	FileSortFieldName FileSortField = "NAME"
	FileSortFieldSlug FileSortField = "SLUG"
)

var AllFileSortField = []FileSortField{
	FileSortFieldName,
	FileSortFieldSlug,
}

func (e FileSortField) IsValid() bool {
	switch e {
	case FileSortFieldName, FileSortFieldSlug:
		return true
	}
	return false
}

func (e FileSortField) String() string {
	return string(e)
}

func (e *FileSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FileSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FileSortField", str)
	}
	return nil
}

func (e FileSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FolderSortField string

const (
	FolderSortFieldName        FolderSortField = "NAME"
	FolderSortFieldSlug        FolderSortField = "SLUG"
	FolderSortFieldDescription FolderSortField = "DESCRIPTION"
)

var AllFolderSortField = []FolderSortField{
	FolderSortFieldName,
	FolderSortFieldSlug,
	FolderSortFieldDescription,
}

func (e FolderSortField) IsValid() bool {
	switch e {
	case FolderSortFieldName, FolderSortFieldSlug, FolderSortFieldDescription:
		return true
	}
	return false
}

func (e FolderSortField) String() string {
	return string(e)
}

func (e *FolderSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FolderSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FolderSortField", str)
	}
	return nil
}

func (e FolderSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GroupSortField string

const (
	GroupSortFieldName        GroupSortField = "NAME"
	GroupSortFieldDescription GroupSortField = "DESCRIPTION"
)

var AllGroupSortField = []GroupSortField{
	GroupSortFieldName,
	GroupSortFieldDescription,
}

func (e GroupSortField) IsValid() bool {
	switch e {
	case GroupSortFieldName, GroupSortFieldDescription:
		return true
	}
	return false
}

func (e GroupSortField) String() string {
	return string(e)
}

func (e *GroupSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GroupSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GroupSortField", str)
	}
	return nil
}

func (e GroupSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PermissionSortField string

const (
	PermissionSortFieldName        PermissionSortField = "NAME"
	PermissionSortFieldDescription PermissionSortField = "DESCRIPTION"
)

var AllPermissionSortField = []PermissionSortField{
	PermissionSortFieldName,
	PermissionSortFieldDescription,
}

func (e PermissionSortField) IsValid() bool {
	switch e {
	case PermissionSortFieldName, PermissionSortFieldDescription:
		return true
	}
	return false
}

func (e PermissionSortField) String() string {
	return string(e)
}

func (e *PermissionSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PermissionSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PermissionSortField", str)
	}
	return nil
}

func (e PermissionSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RoleSortField string

const (
	RoleSortFieldName        RoleSortField = "NAME"
	RoleSortFieldDescription RoleSortField = "DESCRIPTION"
)

var AllRoleSortField = []RoleSortField{
	RoleSortFieldName,
	RoleSortFieldDescription,
}

func (e RoleSortField) IsValid() bool {
	switch e {
	case RoleSortFieldName, RoleSortFieldDescription:
		return true
	}
	return false
}

func (e RoleSortField) String() string {
	return string(e)
}

func (e *RoleSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleSortField", str)
	}
	return nil
}

func (e RoleSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

var AllSortOrder = []SortOrder{
	SortOrderAsc,
	SortOrderDesc,
}

func (e SortOrder) IsValid() bool {
	switch e {
	case SortOrderAsc, SortOrderDesc:
		return true
	}
	return false
}

func (e SortOrder) String() string {
	return string(e)
}

func (e *SortOrder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortOrder", str)
	}
	return nil
}

func (e SortOrder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserSortField string

const (
	UserSortFieldName  UserSortField = "NAME"
	UserSortFieldEmail UserSortField = "EMAIL"
)

var AllUserSortField = []UserSortField{
	UserSortFieldName,
	UserSortFieldEmail,
}

func (e UserSortField) IsValid() bool {
	switch e {
	case UserSortFieldName, UserSortFieldEmail:
		return true
	}
	return false
}

func (e UserSortField) String() string {
	return string(e)
}

func (e *UserSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserSortField", str)
	}
	return nil
}

func (e UserSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
