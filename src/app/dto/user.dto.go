package dto

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6,lte=30"`
	Name     string `json:"name" validate:"required"`
}
