package model

import "github.com/go-playground/validator/v10"

type Folder struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

type CreateFolder struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Slug        string `json:"slug" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,max=500"`
}

type UpdateFolder struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Slug        string `json:"slug" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,max=500"`
}

var validate = validator.New()

// Validate function to validate NewFolder struct
func (f *CreateFolder) Validate() error {
	return validate.Struct(f)
}

// Validate function to validate UpdateFolder struct
func (f *UpdateFolder) Validate() error {
	return validate.Struct(f)
}
