package requests

type InputRegisterUser struct {
	FullName string `json:"fullname" validate:"required,lowercase"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Enabled  bool   `json:"enabled" validate:"required" example:"true"`
}

type InputUpdateUser struct {
	FullName string `json:"fullname" validate:"required,lowercase"`
	Email    string `json:"email" validate:"required,email"`
	Enabled  bool   `json:"enabled" validate:"required" example:"true"`
}

type InputDeleteUser struct {
	ID string `validate:"required,uuid"`
}

type InputGetUser struct {
	ID string `validate:"required,uuid"`
}
