package entities

import (
	"time"

	"gorm.io/gorm"
)

type Collaboration struct {
	PlaylistID uint64 `gorm:"foreignKey"`
	UserID     uint64 `gorm:"foreignKey"`
	CreatedAt  time.Time
}

func (Collaboration) BeforeCreate(db *gorm.DB) (err error) {
	err = db.SetupJoinTable(&Playlist{}, "Users", &Collaboration{})
	if err != nil {

	}
	db.AutoMigrate(&Collaboration{})
	return
}

func ObjCollaboration(playlistID uint64, userID uint64) (collaboration *Collaboration) {
	return &Collaboration{
		PlaylistID: playlistID,
		UserID:     userID,
	}
}
