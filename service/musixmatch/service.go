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
	music, err = s.repository.FindById(id)
	if music == nil || err != nil {
		music, err = s.repository.FindByQuery("musix_id", id)
		if music != nil {
			return
		}
		request, err := http.NewRequest("GET", s.url+"/track.get?apikey="+s.key+"&track_id="+strconv.FormatUint(id, 10), nil)
		if err != nil {
			return nil, err
		}
		var client = &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return nil, goplaylist.ErrInternalServer
		}
		defer response.Body.Close()
		var data dto.Musix
		resBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, goplaylist.ErrInternalServer
		}
		json.Unmarshal(resBody, &data)

		musix := dto.MusixDTO{
			MusixID:    data.Message.Body.Track.MusixID,
			Title:      data.Message.Body.Track.Title,
			Performer:  data.Message.Body.Track.Performer,
			AlbumTitle: data.Message.Body.Track.AlbumTitle,
		}
		err = s.InsertMusix(musix)
		if err != nil {
			return nil, err
		}
		music, err = s.repository.FindByQuery("musix_id", data.Message.Body.Track.MusixID)
	}
	return music, err
}

func (s *service) GetAll() (musics []entities.Music, err error) {
	musics, err = s.repository.FindAll()
	if len(musics) == 0 || err != nil {
		request, err := http.NewRequest("GET", s.url+"/track.search?apikey="+s.key+"&page_size=5&page=1&s_track_rating=desc", nil)
		if err != nil {
			return nil, err
		}
		var client = &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		var data dto.Musix
		resBody, err := ioutil.ReadAll(response.Body)
		// if err != nil {
		// 	return
		// }
		json.Unmarshal(resBody, &data)
		for _, el := range data.Message.Body.TrackList {
			music := dto.MusixDTO{
				MusixID:    el.Track.MusixID,
				Title:      el.Track.Title,
				Performer:  el.Track.Performer,
				AlbumTitle: el.Track.AlbumTitle,
			}
			err = s.InsertMusix(music)
			if err != nil {
				return nil, err
			}
		}
		musics, err = s.repository.FindAll()
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
	newMusic.MusixID = uint64(dto.MusixID)

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
