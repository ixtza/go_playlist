package entities

import (
	"gorm.io/gorm"
)

type Playlist struct {
	ID     uint64 `gorm:"primaryKey"`
	Name   string
	Owner  uint64
	Users  []User  `gorm:"many2many:collaborations"`
	Musics []Music `gorm:"many2many:playlist_musics"`
	gorm.Model
}

func ObjPlaylist(dataName string, dataOwner uint64) (playlist *Playlist) {
	return &Playlist{
		Name:  dataName,
		Owner: dataOwner,
	}
}
