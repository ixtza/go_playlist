package entities

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Playlists []Playlist `gorm:"foreignKey:Owner"`
	gorm.Model
}

func ObjUser(dataName string, dataEmail string, dataPassword string) (user *User) {
	return &User{
		Name:     dataName,
		Email:    dataEmail,
		Password: dataPassword,
	}
}
