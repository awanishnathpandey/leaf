package model

type ChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ForgotPassword struct {
	Email string `json:"email"`
}

type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// type Register struct {
// 	Name     string `json:"name" validate:"required,min=2,max=100"`
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=3,max=50"`
// }

type ResetPassword struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// type AuthUser struct {
// 	ID              int64  `json:"id"`
// 	Name            string `json:"name"`
// 	Email           string `json:"email"`
// 	EmailVerifiedAt *int64 `json:"emailVerifiedAt,omitempty"`
// 	LastSeenAt      int64  `json:"lastSeenAt"`
// 	CreatedAt       int64  `json:"createdAt"`
// 	UpdatedAt       int64  `json:"updatedAt"`
// }

// Validate function to validate NewFolder struct
func (f *Register) Validate() error {
	return validate.Struct(f)
}
