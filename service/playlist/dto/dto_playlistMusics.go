package dto

type PlaylistMusicDTO struct {
	MusicID    uint64 `validate:"required"`
	PlaylistID uint64 `validate:"required"`
}
