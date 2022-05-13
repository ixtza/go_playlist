package collaboration

import (
	goplaylist "mini-clean"
	"mini-clean/entities"
	"mini-clean/service/collaboration/dto"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	// FindById(id uint64) (collaboration *entities.Collaboration, err error)
	// FindAll() (collaborations []entities.Collaboration, err error)
	// FindByQuery(key string, value interface{}) (collaboration entities.Collaboration, err error)
	Exist(userId uint64, playlistId uint64) (collaboration *entities.Collaboration, err error)
	Insert(data entities.Collaboration) (err error)
	Delete(userId uint64, playlistId uint64) (err error)
}

type Service interface {
	// GetById(id uint64) (collaboration *entities.Collaboration, err error)
	// GetAll() (collaborations []entities.Collaboration, err error)
	Exist(userId uint64, playlistId uint64) (result bool, err error)
	Create(dto dto.CollaborationDTO) (err error)
	Remove(userId uint64, playlistId uint64) (result bool, err error)
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

// func (s *service) GetById(id uint64) (collaboration *entities.Collaboration, err error) {
// 	collaboration, err = s.repository.FindById(id)
// 	return
// }

// func (s *service) GetAll() (collaborations []entities.Collaboration, err error) {
// 	collaborations, err = s.repository.FindAll()
// 	return
// }

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

func (s *service) Remove(userId uint64, playlistId uint64) (result bool, err error) {
	err = s.repository.Delete(userId, playlistId)
	if err != nil {
		result = false
	}
	return true, nil
}
