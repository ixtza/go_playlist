package music

import (
	"mini-clean/service/musixMatch"
	"mini-clean/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) musixMatch.Repository {
	var musicRepo musixMatch.Repository

	if dbCon.Driver == util.Postgres {
		musicRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return musicRepo
}
