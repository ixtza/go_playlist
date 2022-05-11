package request

import "mini-clean/service/user/dto"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *CreateUserRequest) ToSpec() *dto.UserDTO {
	return &dto.UserDTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
