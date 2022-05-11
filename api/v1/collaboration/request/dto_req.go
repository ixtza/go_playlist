package request

import "mini-clean/service/collaboration/dto"

type CreateCollborationRequest struct {
	PlaylistID uint64 `json:"playlist_id"`
	UserID     uint64 `json:"user_ud"`
}

func (req *CreateCollborationRequest) ToSpec() *dto.CollaborationDTO {
	return &dto.CollaborationDTO{
		PlaylistID: req.PlaylistID,
		UserID:     req.UserID,
	}
}
