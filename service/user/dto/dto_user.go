package dto

type UserLoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password,omitempty" form:"password" validate:"required"`
}

type UserDTO struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
