package request

import "mini-clean/service/auth/dto"

type AuthRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

func (req *AuthRequest) ToSpec() *dto.InputLogin {
	return &dto.InputLogin{
		Email:    req.Email,
		Password: req.Password,
	}
}
