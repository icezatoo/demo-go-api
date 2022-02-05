package dto

type CreateUserRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
	FullName string `form:"fullName" json:"fullName" binding:"required"`
	LastName string `form:"lastName" json:"lastName" binding:"required"`
	Enabled string `form:"enabled" json:"enabled" binding:"required"`
	Password string `form:"password" json:"password"  binding:"required"`
}