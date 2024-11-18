package model

type File struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	URL       string   `json:"url"`
	FolderID  int64    `json:"folderId"`
	Folder    *Folder  `json:"folder"`
	Groups    []*Group `json:"groups"`
	CreatedAt int64    `json:"createdAt"`
	UpdatedAt int64    `json:"updatedAt"`
}

type UpdateFile struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

type CreateFile struct {
	Name     string `json:"name"`
	Slug     string `json:"slug" validate:"required,min=3,max=100"`
	URL      string `json:"url" validate:"required,min=3"`
	FolderID int64  `json:"folderId"`
}

// Validate function to validate NewFolder struct
func (f *CreateFile) Validate() error {
	return validate.Struct(f)
}
