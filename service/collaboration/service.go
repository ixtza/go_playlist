package collaboration

import (
	"mini-clean/entities"
	goplaylist "mini-clean/error"
	"mini-clean/service/collaboration/dto"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	Exist(userId uint64, playlistId uint64) (collaboration *entities.Collaboration, err error)
	Insert(data entities.Collaboration) (err error)
	Delete(userId uint64, playlistId uint64) (err error)
}

type Service interface {
	Exist(userId uint64, playlistId uint64) (result bool, err error)
	Create(dto dto.CollaborationDTO) (err error)
	Remove(userId uint64, playlistId uint64) (err error)
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

func (s *service) Exist(userId uint64, playlistId uint64) (result bool, err error) {
	_, err = s.repository.Exist(userId, playlistId)
	if err != nil {
		return
	}
	return true, nil
}

func (s *service) Create(dto dto.CollaborationDTO) (err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		return goplaylist.ErrBadRequest
	}

	newCollaboration := entities.ObjCollaboration(dto.PlaylistID, dto.UserID)

	err = s.repository.Insert(*newCollaboration)
	return
}

func (s *service) Remove(userId uint64, playlistId uint64) (err error) {

	err = s.repository.Delete(userId, playlistId)
	if err != nil {
		return
	}
	return
}
