package playlist

import (
	"fmt"
	"mini-clean/entities"
	"mini-clean/service/playlist/dto"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	Exist(id uint64) (playlist *entities.Playlist, err error)
	ExistCollab(userId uint64, playlistId uint64) (playlist *entities.Playlist, err error)
	FindById(id uint64) (playlist *entities.Playlist, err error)
	FindAll() (playlists []entities.Playlist, err error)
	FindByQuery(key string, value interface{}) (playlist []entities.Playlist, err error)
	Insert(data entities.Playlist) (err error)
	Update(data entities.Playlist) (playlist *entities.Playlist, err error)
	Delete(id uint64) (err error)
	AddPlaylistMusic(data entities.PlaylistMusic) (err error)
	FindPlaylistMusicById(playlistId uint64) (playlistMusics entities.Playlist, err error)
	DeletePlaylistMusicById(musicId uint64, playlistId uint64) (err error)
}

type Service interface {
	Ownership(userId uint64, playlistId uint64) (err error)
	Access(userId uint64, playlistId uint64) (err error)
	GetById(id uint64) (playlist *entities.Playlist, err error)
	GetAll() (playlists []entities.Playlist, err error)
	Create(dto dto.PlaylistDTO) (err error)
	Modify(id uint64, dto dto.PlaylistDTO) (playlist *entities.Playlist, err error)
	Remove(userId uint64, playlistId uint64) (err error)
	AddPlaylistMusic(userId uint64, dto dto.PlaylistMusicDTO) (err error)
	GetPlaylistMusicById(userId uint64, playlistId uint64) (playlist entities.Playlist, err error)
	RemovePlaylistMusicById(userId uint64, musicId uint64, playlistId uint64) (err error)
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

func (s *service) Ownership(userId uint64, playlistId uint64) (err error) {
	_, err = s.repository.FindByQuery("owner", userId)
	// internal server error / error not found
	if err != nil {
		return
	}
	return
}

func (s *service) Access(userId uint64, playlistId uint64) (err error) {
	_, err = s.repository.ExistCollab(userId, playlistId)
	if err != nil {
		return
	}
	return
}

func (s *service) GetById(id uint64) (playlist *entities.Playlist, err error) {
	playlist, err = s.repository.FindById(id)
	// internal server error / error not found
	return
}

func (s *service) GetAll() (playlists []entities.Playlist, err error) {
	playlists, err = s.repository.FindAll()
	// internal server error / error not found
	return
}

func (s *service) Create(dto dto.PlaylistDTO) (err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		// bad request
		return err
	}

	newPlaylist := entities.ObjPlaylist(dto.Name, dto.Owner)

	err = s.repository.Insert(*newPlaylist)
	// internal server error
	return
}

func (s *service) Modify(id uint64, dto dto.PlaylistDTO) (playlist *entities.Playlist, err error) {
	err = s.validate.Struct(dto)
	if err != nil {
		// bad request
		return nil, err
	}
	playlist, err = s.repository.FindById(id)
	// internal server error / error not found
	if err != nil {
		return nil, err
	}

	playlist.Name = dto.Name

	playlist, err = s.repository.Update(*playlist)
	if err != nil {
		// internal server error / error not found
		return nil, err
	}
	return playlist, nil
}

func (s *service) Remove(userId uint64, playlistId uint64) (err error) {
	err = s.Ownership(userId, playlistId)
	// error unauthorized
	if err != nil {
		return
	}
	err = s.repository.Delete(playlistId)
	return
}

func (s *service) AddPlaylistMusic(userId uint64, dto dto.PlaylistMusicDTO) (err error) {
	err = s.Access(userId, dto.PlaylistID)
	if err != nil {
		err = s.Ownership(userId, dto.PlaylistID)
		// error unauthorized / error not found
		if err != nil {
			return
		}
		newPlaylistMusic := entities.ObjPlaylistMusics(dto.MusicID, dto.PlaylistID)
		err = s.repository.AddPlaylistMusic(*newPlaylistMusic)
		// internal server error / error not found
		if err != nil {
			return
		}
		return
	}
	newPlaylistMusic := entities.ObjPlaylistMusics(dto.MusicID, dto.PlaylistID)
	err = s.repository.AddPlaylistMusic(*newPlaylistMusic)
	// internal server error / error not found
	if err != nil {
		return
	}
	return
}

func (s *service) GetPlaylistMusicById(userId uint64, playlistId uint64) (playlist entities.Playlist, err error) {
	playlist, err = s.repository.FindPlaylistMusicById(playlistId)
	// internal server error / error not found
	if err != nil {
		return
	}
	return
}

func (s *service) RemovePlaylistMusicById(userId uint64, musicId uint64, playlistId uint64) (err error) {
	err = s.Access(userId, playlistId)
	fmt.Println(err)
	if err != nil {
		err = s.Ownership(userId, playlistId)
		// error unauthorized / error not found
		if err != nil {
			return
		}
		err = s.repository.DeletePlaylistMusicById(musicId, playlistId)
		// internal server error / error not found
		if err != nil {
			return
		}
		return
	}
	err = s.repository.DeletePlaylistMusicById(musicId, playlistId)
	// internal server error / error not found
	if err != nil {
		return
	}
	return
}
