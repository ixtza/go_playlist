package musixmatch

import (
	"encoding/json"
	"io/ioutil"
	goplaylist "mini-clean"
	"mini-clean/entities"
	"mini-clean/service/musixmatch/dto"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id uint64) (music *entities.Music, err error)
	FindAll() (musics []entities.Music, err error)
	FindByQuery(key string, value interface{}) (music *entities.Music, err error)
	Insert(data entities.Music) (id uint64, err error)
	Update(data entities.Music) (music *entities.Music, err error)
	Delete(id uint64) (err error)
}

type Service interface {
	GetById(id uint64) (music *entities.Music, err error)
	GetAll() (musics []entities.Music, err error)
	Create(dto dto.MusicDTO) (id uint64, err error)
	Modify(id uint64, dto dto.MusicDTO) (music *entities.Music, err error)
	Remove(id uint64) (err error)
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
	music, err = s.repository.FindById(id)

	if err != nil {
		if id < 100000000 {
			id += 100000000
		}

		music, err = s.repository.FindByQuery("musix_id", id)
		if err != nil {

			request, err := http.NewRequest("GET", s.url+"/track.get?apikey="+s.key+"&track_id="+strconv.FormatUint(id, 10), nil)
			var client = &http.Client{}
			response, err := client.Do(request)
			var data dto.Musix
			resBody, err := ioutil.ReadAll(response.Body)
			err = json.Unmarshal(resBody, &data)

			musix := dto.MusicDTO{
				MusixID:    data.Message.Body.Track.MusixID,
				Title:      data.Message.Body.Track.Title,
				Performer:  data.Message.Body.Track.Performer,
				AlbumTitle: data.Message.Body.Track.AlbumTitle,
			}
			_, err = s.Create(musix)
			music, err = s.repository.FindByQuery("musix_id", data.Message.Body.Track.MusixID)
			return music, err
		}
	}
	return
}

func (s *service) GetAll() (musics []entities.Music, err error) {
	musics, err = s.repository.FindAll()
	if len(musics) == 0 || err != nil {
		var request *http.Request
		request, err = http.NewRequest("GET", s.url+"/track.search?apikey="+s.key+"&page_size=5&page=1&s_track_rating=desc", nil)
		var client = &http.Client{}
		response, err := client.Do(request)
		var data dto.Musix
		resBody, err := ioutil.ReadAll(response.Body)
		json.Unmarshal(resBody, &data)
		for _, el := range data.Message.Body.TrackList {
			music := dto.MusicDTO{
				MusixID:    el.Track.MusixID,
				Title:      el.Track.Title,
				Performer:  el.Track.Performer,
				AlbumTitle: el.Track.AlbumTitle,
			}
			_, err = s.Create(music)
		}
		if err != nil {
			return nil, goplaylist.ErrInternalServer
		}
		musics, err = s.repository.FindAll()
	}
	return musics, nil
}

func (s *service) Create(dto dto.MusicDTO) (id uint64, err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		err = goplaylist.ErrBadRequest
		return
	}

	newMusic := entities.ObjMusic(dto.Title, dto.Performer, dto.AlbumTitle)
	newMusic.MusixID = dto.MusixID
	id, err = s.repository.Insert(*newMusic)
	return
}

func (s *service) Modify(id uint64, dto dto.MusicDTO) (music *entities.Music, err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		return nil, goplaylist.ErrBadRequest
	}
	music, err = s.repository.FindById(id)
	music.Title = dto.Title
	music.Performer = dto.Performer
	music, err = s.repository.Update(*music)
	if err != nil {
		return
	}
	return
}

func (s *service) Remove(id uint64) (err error) {
	err = s.repository.Delete(id)
	if err != nil {
		return
	}
	return
}
