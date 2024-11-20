package model

type CreateGroup struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=100"`
}

type Group struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CreatedAt   int64           `json:"createdAt"`
	UpdatedAt   int64           `json:"updatedAt"`
	Users       *UserConnection `json:"users"`
	Folders     []*Folder       `json:"folders"`
	Files       *FileConnection `json:"files"`
}

type UpdateGroup struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=100"`
}

// Validate function to validate NewGroup struct
func (f *CreateGroup) Validate() error {
	return validate.Struct(f)
}

// Validate function to validate UpdateGroup struct
func (f *UpdateGroup) Validate() error {
	return validate.Struct(f)
}
