package modules

import (
	"mini-clean/api"
	"mini-clean/config"
	"mini-clean/util"

	"mini-clean/api/middleware"
	userV1Controller "mini-clean/api/v1/user"
	userRepo "mini-clean/repository/user"
	userService "mini-clean/service/user"

	playlistV1Controller "mini-clean/api/v1/playlist"
	playlistRepo "mini-clean/repository/playlist"
	playlistService "mini-clean/service/playlist"

	musicV1Controller "mini-clean/api/v1/music"
	musicRepo "mini-clean/repository/music"

	// musicService "mini-clean/service/music"

	musixMatchService "mini-clean/service/musixMatch"

	collaborationV1Controller "mini-clean/api/v1/collaboration"
	collaborationRepo "mini-clean/repository/collaboration"
	collaborationService "mini-clean/service/collaboration"

	authController "mini-clean/api/v1/auth"
	authService "mini-clean/service/auth"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userPermitRepository := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewService(userPermitRepository)
	userV1PermitController := userV1Controller.NewController(userPermitService)

	musicPermitRepository := musicRepo.RepositoryFactory(dbCon)
	musixMatchPermitService := musixMatchService.NewService(musicPermitRepository, config.OpenApi.MusixMatch, config.OpenApi.MusixMatchUrl)
	musicV1PermitController := musicV1Controller.NewController(musixMatchPermitService)

	playlistPermitRepository := playlistRepo.RepositoryFactory(dbCon)
	playlistPermitService := playlistService.NewService(playlistPermitRepository)
	playlistV1PermitController := playlistV1Controller.NewController(playlistPermitService)

	collaborationPermitRepository := collaborationRepo.RepositoryFactory(dbCon)
	collaborationPermitService := collaborationService.NewService(collaborationPermitRepository)
	collaborationV1PermitController := collaborationV1Controller.NewController(collaborationPermitService, playlistPermitService)

	middlewarePermitJwt := middleware.NewJwtService(config.App.JWTKey)

	authPermitService := authService.NewService(config, userPermitRepository)
	authPermitController := authController.NewController(authPermitService)

	controllers := api.Controller{
		UserV1Controller:     userV1PermitController,
		PlaylistV1Controller: playlistV1PermitController,
		MusicV1Controller:    musicV1PermitController,
		CollabV1Controller:   collaborationV1PermitController,
		AuthV1Controller:     authPermitController,
		MiddlewareJwt:        middlewarePermitJwt,
	}

	return controllers
}
