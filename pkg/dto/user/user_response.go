package dto

type UserReponse struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Enabled       bool `json:"enabled"`
}


