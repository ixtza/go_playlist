package auth

import (
	"mini-clean/config"
	"mini-clean/entities"
	goplaylist "mini-clean/error"
	"mini-clean/service/auth/dto"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type Repository interface {
	FindByQuery(key string, value interface{}) (user *entities.User, err error)
}

type Service interface {
	Login(input dto.InputLogin) (auth *entities.Auth, err error)
}

type service struct {
	validate   *validator.Validate
	config     *config.AppConfig
	repository Repository
}

// NewService ...
func NewService(config *config.AppConfig, repository Repository) Service {
	return &service{
		validate:   validator.New(),
		config:     config,
		repository: repository,
	}
}

func (s *service) Login(input dto.InputLogin) (auth *entities.Auth, err error) {
	err = s.validate.Struct(input)
	if err != nil {
		return nil, goplaylist.ErrBadRequest
	}

	data := new(entities.User)
	data, err = s.repository.FindByQuery("email", input.Email)

	if err != nil {
		return nil, err
	}

	excepctedPassword := data.Password
	if excepctedPassword != input.Password {
		return nil, goplaylist.ErrUnauthorized
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &entities.Claims{
		ID:       data.ID,
		Username: input.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	screetKey := []byte(s.config.App.JWTKey)

	tokenString, err := token.SignedString(screetKey)
	if err != nil {
		return nil, goplaylist.ErrInternalServer
	}
	auth = &entities.Auth{
		Token: tokenString,
	}

	return auth, nil
}
