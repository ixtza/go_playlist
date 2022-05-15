package playlist

import (
	v1 "mini-clean/api/v1"
	"mini-clean/api/v1/playlist/request"
	"mini-clean/api/v1/playlist/response"
	playlistUsecase "mini-clean/service/playlist"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service playlistUsecase.Service
}

func NewController(service playlistUsecase.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// GetByID godoc
// @Summary      Get playlist
// @Description  Retrive playlist data by id
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "Playlist ID"
// @Success      200  {object}  response.PlaylistResponseSuccess
// @Failure      400  {object}  response.PlaylistResponse
// @Failure      403  {object}  response.PlaylistResponse
// @Failure      500  {object}  response.PlaylistResponse
// @Security ApiKeyAuth
// @Router       /playlists/{id} [get]
func (controller *Controller) GetByID(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	playlist, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponseSuccess{
		Status: "success",
		Data:   playlist,
	})
}

// Modify godoc
// @Summary      Modify playlist
// @Description  Edit playlist data by id
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "Playlist ID"
// @Success      200  {object}  response.PlaylistResponse
// @Failure      400  {object}  response.PlaylistResponse
// @Failure      401  {object}  response.PlaylistResponse
// @Failure      403  {object}  response.PlaylistResponse
// @Failure      500  {object}  response.PlaylistResponse
// @Security ApiKeyAuth
// @Router       /playlists/{id} [put]
func (controller *Controller) Modify(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	createPlaylistRequest := new(request.CreatePlaylistRequest)
	if err := c.Bind(createPlaylistRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	err = controller.service.Ownership(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createPlaylistRequest.ToSpec(uint64(userId))
	_, err = controller.service.Modify(uint64(id), req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
		Status:  "success",
		Message: "playlist " + params + " updated",
	})
}

// Create godoc
// @Summary      Create playlist
// @Description  Make new playlist
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param Payload body request.CreatePlaylistRequest true "Playlist json format" SchemaExample(request.CreatePlaylistRequest)
// @Success      201  {object}  response.PlaylistResponse
// @Failure      400  {object}  response.PlaylistResponse
// @Failure      403  {object}  response.PlaylistResponse
// @Failure      500  {object}  response.PlaylistResponse
// @Security ApiKeyAuth
// @Router       /playlists [post]
func (controller *Controller) Create(c echo.Context) error {

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}

	createPlaylistRequest := new(request.CreatePlaylistRequest)
	if err := c.Bind(createPlaylistRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createPlaylistRequest.ToSpec(uint64(userId))
	id, err := controller.service.Create(req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.PlaylistResponse{
		Status:  "success",
		Message: map[string]interface{}{"playlist_id": id},
	})
}

// Delete godoc
// @Summary      Delete playlist
// @Description  Remove playlist by id
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "Playlist ID"
// @Success      200  {object}  response.PlaylistResponse
// @Failure      400  {object}  response.PlaylistResponse
// @Failure      403  {object}  response.PlaylistResponse
// @Failure      500  {object}  response.PlaylistResponse
// @Security ApiKeyAuth
// @Router       /playlists/{id} [delete]
func (controller *Controller) Delete(c echo.Context) error {

	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	err = controller.service.Ownership(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	err = controller.service.Remove(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
		Status:  "success",
		Message: "playlist id's " + params + " deleted",
	})
}

// GetAll godoc
// @Summary      Get all playlist
// @Description  Retrive all playlist data
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.PlaylistResponseSuccess
// @Failure      400  {object}  response.PlaylistResponse
// @Failure      403  {object}  response.PlaylistResponse
// @Failure      500  {object}  response.PlaylistResponse
// @Security ApiKeyAuth
// @Router       /playlists [get]
func (controller *Controller) GetAll(c echo.Context) error {
	playlists, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponseSuccess{
		Status: "success",
		Data:   playlists,
	})
}

// AddPlaylistMusic godoc
// @Summary      Add playlist music
// @Description  Add music to playlist
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "Playlist ID"
// @Param Payload body request.CreatePlaylistRequest true "Playlist json format" SchemaExample(request.CreatePlaylistRequest)
// @Success      200  {object}  response.PlaylistResponse
// @Failure      400  {object}  response.PlaylistResponse
// @Failure      401  {object}  response.PlaylistResponse
// @Failure      403  {object}  response.PlaylistResponse
// @Failure      500  {object}  response.PlaylistResponse
// @Security ApiKeyAuth
// @Router       /playlists/{id}/songs [post]
func (controller *Controller) AddPlaylistMusic(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	createPlaylistMusicRequest := new(request.CreatePlaylistMusicRequset)
	if err := c.Bind(createPlaylistMusicRequest); err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createPlaylistMusicRequest.ToSpec(uint64(id))

	err = controller.service.AddPlaylistMusic(uint64(userId), req)

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
		Status:  "success",
		Message: "music added to playlist",
	})
}

// GetPlaylistMusicById godoc
// @Summary      Get playlist's music
// @Description  Retrive playlist's music data by id
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "Playlist ID"
// @Success      200  {object}  response.PlaylistResponseSuccess
// @Failure      400  {object}  response.PlaylistResponse
// @Failure      403  {object}  response.PlaylistResponse
// @Failure      404  {object}  response.PlaylistResponse
// @Failure      500  {object}  response.PlaylistResponse
// @Security ApiKeyAuth
// @Router       /playlists/{id}/songs [get]
func (controller *Controller) GetPlaylistMusicById(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	res, err := controller.service.GetPlaylistMusicById(uint64(userId), uint64(id))
	if err != nil {

		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponseSuccess{
		Status: "success",
		Data:   map[string]interface{}{"playlist_id": id, "musics_list": res.Musics},
	})
}

// RemovePlaylistMusicById godoc
// @Summary      Remove playlist's music
// @Description  Delete playlist's music from playlist
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "Playlist ID"
// @Param Payload body request.CreatePlaylistRequest true "Playlist json format" SchemaExample(request.CreatePlaylistRequest)
// @Success      200  {object}  response.PlaylistResponse
// @Failure      400  {object}  response.PlaylistResponse
// @Failure      401  {object}  response.PlaylistResponse
// @Failure      403  {object}  response.PlaylistResponse
// @Failure      500  {object}  response.PlaylistResponse
// @Security ApiKeyAuth
// @Router       /playlists/{id}/songs [delete]
func (controller *Controller) RemovePlaylistMusicById(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	createPlaylistMusicRequest := new(request.CreatePlaylistMusicRequset)
	if err := c.Bind(createPlaylistMusicRequest); err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	err = controller.service.RemovePlaylistMusicById(uint64(userId), createPlaylistMusicRequest.MusicID, uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.PlaylistResponse{
		Status:  "success",
		Message: "playlist id's " + params + " removed",
	})
}
