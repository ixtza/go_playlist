package user

import (
	"fmt"
	goplaylist "mini-clean"
	"mini-clean/entities"
	"mini-clean/service/user/dto"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id uint64) (user *entities.User, err error)
	FindAll() (users []entities.User, err error)
	FindByQuery(key string, value interface{}) (user *entities.User, err error)
	Insert(data entities.User) (id uint64, err error)
	Update(data entities.User) (user *entities.User, err error)
}

type Service interface {
	GetById(id uint64) (user *entities.User, err error)
	GetAll() (users []entities.User, err error)
	Create(dto dto.UserDTO) (id uint64, err error)
	Modify(dto dto.UserDTO) (user *entities.User, err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) GetById(id uint64) (user *entities.User, err error) {
	result, err := s.repository.FindById(id)
	return result, err
}

func (s *service) GetAll() (users []entities.User, err error) {
	users, err = s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Create(dto dto.UserDTO) (id uint64, err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		err = goplaylist.ErrBadRequest
		return
	}

	newUser := entities.ObjUser(dto.Name, dto.Email, dto.Password)

	id, err = s.repository.Insert(*newUser)
	return
}

func (s *service) Modify(dto dto.UserDTO) (user *entities.User, err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		return nil, goplaylist.ErrBadRequest
	}

	user, err = s.repository.FindByQuery("email", dto.Email)
	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	user.Name = dto.Name
	user.Password = dto.Password

	user, err = s.repository.Update(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
