package request

import "mini-clean/service/musixmatch/dto"

type CreateMusicRequest struct {
	Title      string `json:"title"`
	Performer  string `json:"performer"`
	MusixID    uint64 `json:"musix_id"`
	AlbumTitle string `json:"album_title"`
}

func (req *CreateMusicRequest) ToSpec() *dto.MusicDTO {
	return &dto.MusicDTO{
		Title:      req.Title,
		Performer:  req.Performer,
		MusixID:    req.MusixID,
		AlbumTitle: req.AlbumTitle,
	}
}
