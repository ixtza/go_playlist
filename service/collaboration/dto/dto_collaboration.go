package dto

type CollaborationDTO struct {
	PlaylistID uint64 `validate:"required"`
	UserID     uint64 `validate:"required"`
}
