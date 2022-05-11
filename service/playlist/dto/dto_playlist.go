package dto

type PlaylistDTO struct {
	Name  string `validate:"required"`
	Owner uint64 `validate:"required"`
}
