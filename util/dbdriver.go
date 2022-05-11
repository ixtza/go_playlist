package util

import (
	"mini-clean/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	Postgres DatabaseDriver = "postgres"

	Static DatabaseDriver = "static"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	Postgres *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Database.Driver {
	case "postgres":
		db.Driver = Postgres
		db.Postgres = newPostgres(config)
	case "static":
		db.Driver = Static
	default:
		panic("unsupported driver")
	}

	return &db
}

func newPostgres(config *config.AppConfig) *gorm.DB {
	dbURL := config.Database.DBURL
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.Migrator().DropTable(&entities.User{}, &entities.Music{}, &entities.Playlist{}, &entities.Collaboration{}, &entities.PlaylistMusic{})
	// db.AutoMigrate(&entities.User{}, &entities.Music{}, &entities.Playlist{})
	// entities.Collaboration{}.BeforeCreate(db)
	// entities.PlaylistMusic{}.BeforeCreate(db)

	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.Postgres != nil {
		db, _ := db.Postgres.DB()
		db.Close()
	}
}
