package util

import (
	"fmt"
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
	var connectionString string
	switch config.Database.Driver {
	case "postgres":
		connectionString = fmt.Sprintf("%s://%s:%s@%s:%s/%s",
			config.Database.Driver,
			config.Database.DB_USER,
			config.Database.DB_PASSWORD,
			config.Database.DB_HOST,
			config.Database.DB_PORT,
			config.Database.DB_NAME)
	}
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.Postgres != nil {
		db, _ := db.Postgres.DB()
		db.Close()
	}
}
