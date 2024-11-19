package model

type Role struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CreatedAt   int64         `json:"createdAt"`
	UpdatedAt   int64         `json:"updatedAt"`
	Permissions []*Permission `json:"permissions"`
	Users       []*User       `json:"users"`
}

type UpdatePermission struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateRole struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Permission struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
	Roles       []*Role `json:"roles"`
}

type CreatePermission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRole struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate function to validate NewGroup struct
func (f *CreateRole) Validate() error {
	return validate.Struct(f)
}

// Validate function to validate UpdateGroup struct
func (f *UpdateRole) Validate() error {
	return validate.Struct(f)
}

// Validate function to validate NewGroup struct
func (f *CreatePermission) Validate() error {
	return validate.Struct(f)
}

// Validate function to validate UpdateGroup struct
func (f *UpdatePermission) Validate() error {
	return validate.Struct(f)
}
