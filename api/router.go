package api

import (
	"mini-clean/api/middleware"
	"mini-clean/api/v1/auth"
	"mini-clean/api/v1/collaboration"
	"mini-clean/api/v1/music"
	"mini-clean/api/v1/playlist"
	"mini-clean/api/v1/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserV1Controller     *user.Controller
	PlaylistV1Controller *playlist.Controller
	MusicV1Controller    *music.Controller
	CollabV1Controller   *collaboration.Controller
	AuthV1Controller     *auth.Controller
	MiddlewareJwt        middleware.JWTService
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.POST("/v1/signup", controller.UserV1Controller.Create)

	authV1 := e.Group("/v1/auth")
	authV1.POST("", controller.AuthV1Controller.Auth)

	userV1 := e.Group("/v1/users")
	userV1.Use(controller.MiddlewareJwt.JwtMiddleware())
	userV1.GET("/:id", controller.UserV1Controller.GetByID)
	userV1.PUT("", controller.UserV1Controller.Modify)
	// userV1.DELETE("/:id", controller.UserV1Controller.Delete)

	playlistV1 := e.Group("/v1/playlists")
	playlistV1.Use(controller.MiddlewareJwt.JwtMiddleware())
	playlistV1.POST("", controller.PlaylistV1Controller.Create)
	playlistV1.GET("", controller.PlaylistV1Controller.GetAll)
	playlistV1.GET("/:id", controller.PlaylistV1Controller.GetByID)
	playlistV1.PUT("/:id", controller.PlaylistV1Controller.Modify)
	playlistV1.DELETE("/:id", controller.PlaylistV1Controller.Delete)

	playlistV1.POST("/:id/songs", controller.PlaylistV1Controller.AddPlaylistMusic)
	playlistV1.GET("/:id/songs", controller.PlaylistV1Controller.GetPlaylistMusicById)
	playlistV1.DELETE("/:id/songs", controller.PlaylistV1Controller.RemovePlaylistMusicById)

	musicV1 := e.Group("/v1/musics")
	musicV1.Use(controller.MiddlewareJwt.JwtMiddleware())
	musicV1.GET("", controller.MusicV1Controller.GetAll)
	musicV1.GET("/:id", controller.MusicV1Controller.GetByID)
	musicV1.PUT("/:id", controller.MusicV1Controller.Modify)
	musicV1.DELETE("/:id", controller.MusicV1Controller.Delete)

	collaborationV1 := e.Group("/v1/collaborations")
	collaborationV1.Use(controller.MiddlewareJwt.JwtMiddleware())
	collaborationV1.POST("", controller.CollabV1Controller.Create)
	collaborationV1.DELETE("", controller.CollabV1Controller.Remove)

	// playlistMusicsV1 := e.Group("/v1/playlistmusics")
	// playlistMusicsV1.Use(controller.MiddlewareJwt.JwtMiddleware())
	// playlistMusicsV1.POST("", controller.PlaylistMusicsV1Controller.Create)
	// playlistMusicsV1.GET("/:id",controller.PlaylistMusicsV1Controller)
	// playlistMusicsV1.DELETE("/:id",controller.PlaylistMusicsV1Controller)

	// admin := e.Group("/v1/admin")
	// admin.Use(controller.MiddlewareJwt.JwtAdminMiddleware())
	// admin := e.GET("/users", controller.UserV1Controller.GetAll)
}
