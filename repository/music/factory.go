package music

import (
	"mini-clean/service/musixmatch"
	"mini-clean/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) musixmatch.Repository {
	var musicRepo musixmatch.Repository

	if dbCon.Driver == util.Postgres {
		musicRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return musicRepo
}
