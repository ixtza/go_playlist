package musixmatch

import (
	"encoding/json"
	"mini-clean/entities"
	"mini-clean/service/music/dto"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id uint64) (music *entities.Music, err error)
	FindAll() (musics []entities.Music, err error)
	FindByQuery(key string, value interface{}) (music entities.Music, err error)
	Insert(data entities.Music) (err error)
	Update(data entities.Music) (music *entities.Music, err error)
	Delete(id uint64) (err error)
}

type Service interface {
	GetById(id uint64) (music *entities.Music, err error)
	GetAll() (musics []entities.Music, err error)
	Create(dto dto.MusicDTO) (err error)
	Modify(id uint64, dto dto.MusicDTO) (music *entities.Music, err error)
	Remove(id uint64) (result bool, err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
	key        string
	url        string
}

func NewService(repository Repository, key string, url string) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
		key:        key,
		url:        url,
	}
}

func (s *service) GetById(id uint64) (music *entities.Music, err error) {
	result, err := s.repository.FindById(id)
	if err != nil || result == nil {
		data := dto.MusixDTO{}
		request, err := http.NewRequest("POST", s.url+"/track.get?track_id="+strconv.FormatUint(id, 10), nil)
		if err != nil {
			return nil, err
		}
		var client = &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return nil, err
		}

		err = s.InsertMusix(data)
		if err != nil {
			return nil, err
		}
	}
	return result, err
}

func (s *service) GetAll() (musics []entities.Music, err error) {
	musics, err = s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return musics, nil
}

func (s *service) Create(dto dto.MusicDTO) (err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		return err
	}

	newMusic := entities.ObjMusic(dto.Title, dto.Performer, dto.AlbumTitle)

	err = s.repository.Insert(*newMusic)
	return err
}

func (s *service) InsertMusix(dto dto.MusixDTO) (err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		return err
	}

	newMusic := entities.ObjMusic(dto.Title, dto.Performer, dto.AlbumTitle)
	newMusic.MusixID = dto.MusixID

	err = s.repository.Insert(*newMusic)
	return err
}

func (s *service) Modify(id uint64, dto dto.MusicDTO) (music *entities.Music, err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		return nil, err
	}
	music, err = s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	music.Title = dto.Title
	music.Performer = dto.Performer

	music, err = s.repository.Update(*music)
	if err != nil {
		return nil, err
	}
	return music, nil
}

func (s *service) Remove(id uint64) (result bool, err error) {
	err = s.repository.Delete(id)
	if err != nil {
		result = false
	}
	return
}