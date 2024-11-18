package model

import "github.com/go-playground/validator/v10"

type User struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	EmailVerifiedAt *int64   `json:"emailVerifiedAt,omitempty"`
	LastSeenAt      int64    `json:"lastSeenAt"`
	CreatedAt       int64    `json:"createdAt"`
	UpdatedAt       int64    `json:"updatedAt"`
	DeletedAt       *int64   `json:"deletedAt,omitempty"`
	Groups          []*Group `json:"groups"`
}

type CreateUser struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=50"`
}

type UpdateUser struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var validate = validator.New()

// Validate function to validate NewFolder struct
func (f *CreateUser) Validate() error {
	return validate.Struct(f)
}

// Validate function to validate UpdateFolder struct
func (f *UpdateUser) Validate() error {
	return validate.Struct(f)
}
