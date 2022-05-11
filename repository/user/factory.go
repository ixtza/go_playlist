package user

import (
	"mini-clean/service/user"
	"mini-clean/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) user.Repository {
	var userRepo user.Repository

	if dbCon.Driver == util.Postgres {
		userRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return userRepo
}
