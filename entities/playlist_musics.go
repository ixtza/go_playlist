package entities

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PlaylistMusic struct {
	MusicID    uint64 `gorm:"foreignKey"`
	PlaylistID uint64 `gorm:"foreignKey"`
	CreatedAt  time.Time
}

func (PlaylistMusic) BeforeCreate(db *gorm.DB) (err error) {
	err = db.SetupJoinTable(&Playlist{}, "Musics", &PlaylistMusic{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&PlaylistMusic{})
	return
}

func ObjPlaylistMusics(musicId uint64, playlistId uint64) (playlistMusics *PlaylistMusic) {
	return &PlaylistMusic{
		MusicID:    musicId,
		PlaylistID: playlistId,
	}
}
