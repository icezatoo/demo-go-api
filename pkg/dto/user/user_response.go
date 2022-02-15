package dto

type UserReponse struct {
	ID       string `copier:"must" json:"id"`
	FullName string `copier:"must" json:"fullName"`
	LastName string `copier:"must" json:"lastName"`
	Email    string `copier:"must" json:"email"`
	Enabled  bool   `copier:"must" json:"enabled"`
}
