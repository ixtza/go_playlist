package dto

type MusicDTO struct {
	Title      string `validate:"required"`
	Performer  string `validate:"required"`
	AlbumTitle string
	MusixID    uint64
}

type MusixDTO struct {
	Title      string `json:"track_name"`
	Performer  string `json:"artist_name"`
	AlbumTitle string `json:"album_name"`
	MusixID    uint64 `json:"track_id"`
}

type TrackList struct {
	Track MusixDTO `json:"track"`
}

type Body struct {
	TrackList []TrackList `json:"track_list"`
	Track     MusixDTO    `json:"track"`
}

type Message struct {
	Body Body `json:"body"`
}

type Musix struct {
	Message Message `json:"message"`
}
