package music

import (
	"mini-clean/service/music"
	"mini-clean/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) music.Repository {
	var musicRepo music.Repository

	if dbCon.Driver == util.Postgres {
		musicRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return musicRepo
}
