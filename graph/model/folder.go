package model

// type Folder struct {
// 	ID          int64    `json:"id"`
// 	Name        string   `json:"name"`
// 	Slug        string   `json:"slug"`
// 	Description string   `json:"description"`
// 	CreatedAt   int64    `json:"createdAt"`
// 	UpdatedAt   int64    `json:"updatedAt"`
// 	Groups      []*Group `json:"groups"`
// 	Files       []*File  `json:"files"`
// }

// type CreateFolder struct {
// 	Name        string `json:"name" validate:"required,min=3,max=100"`
// 	Slug        string `json:"slug" validate:"required,min=3,max=100"`
// 	Description string `json:"description" validate:"required,max=500"`
// }

// type UpdateFolder struct {
// 	ID          int64  `json:"id"`
// 	Name        string `json:"name" validate:"required,min=3,max=100"`
// 	Slug        string `json:"slug" validate:"required,min=3,max=100"`
// 	Description string `json:"description" validate:"required,max=500"`
// }

// Validate function to validate NewFolder struct
func (f *CreateFolder) Validate() error {
	return validate.Struct(f)
}

// Validate function to validate UpdateFolder struct
func (f *UpdateFolder) Validate() error {
	return validate.Struct(f)
}
