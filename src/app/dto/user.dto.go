package dto

type UserDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6,lte=30"`
	Username string `json:"username" validate:"required"`
}

type UpdateUserDto struct {
	Password string `json:"password" validate:"required,gte=6,lte=30"`
	Username string `json:"username" validate:"required"`
}
