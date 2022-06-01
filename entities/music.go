package entities

import "gorm.io/gorm"

type Music struct {
	ID         uint64 `gorm:"primaryKey"`
	MusixID    uint64
	Title      string
	Performer  string
	AlbumTitle string
	gorm.Model
}

func ObjMusic(dataTitle string, dataPerformer string, dataAlbumTitle string, dataMusix uint64) (music *Music) {
	return &Music{
		Title:      dataTitle,
		Performer:  dataPerformer,
		AlbumTitle: dataAlbumTitle,
		MusixID:    dataMusix,
	}
}
