package collaboration

import (
	"mini-clean/service/collaboration"
	"mini-clean/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) collaboration.Repository {
	var collaborationRepo collaboration.Repository

	if dbCon.Driver == util.Postgres {
		collaborationRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return collaborationRepo
}
