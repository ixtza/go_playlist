package request

import "mini-clean/service/playlist/dto"

type CreatePlaylistRequest struct {
	Name  string `json:"name"`
	Owner uint64
}

func (req *CreatePlaylistRequest) ToSpec(ownerId uint64) *dto.PlaylistDTO {
	return &dto.PlaylistDTO{
		Name:  req.Name,
		Owner: ownerId,
	}
}

type CreatePlaylistMusicRequset struct {
	MusicID uint64 `json:"music_id"`
}

func (req *CreatePlaylistMusicRequset) ToSpec(playlistID uint64) *dto.PlaylistMusicDTO {
	return &dto.PlaylistMusicDTO{
		MusicID:    req.MusicID,
		PlaylistID: playlistID,
	}
}
