package playlist

import (
	"mini-clean/service/playlist"
	"mini-clean/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) playlist.Repository {
	var playlistRepo playlist.Repository

	if dbCon.Driver == util.Postgres {
		playlistRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return playlistRepo
}
