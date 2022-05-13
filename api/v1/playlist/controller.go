package playlist

import (
	"fmt"
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

func (controller *Controller) GetByID(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	playlist, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.ResponseSuccess{
		Status: "success",
		Data:   playlist,
	})
}

func (controller *Controller) Modify(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	createPlaylistRequest := new(request.CreatePlaylistRequest)
	if err := c.Bind(createPlaylistRequest); err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	err = controller.service.Ownership(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	req := *createPlaylistRequest.ToSpec(uint64(userId))
	_, err = controller.service.Modify(uint64(id), req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.Response{
		Message: "success",
	})
}

func (controller *Controller) Create(c echo.Context) error {

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}

	createPlaylistRequest := new(request.CreatePlaylistRequest)
	if err := c.Bind(createPlaylistRequest); err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	req := *createPlaylistRequest.ToSpec(uint64(userId))
	err = controller.service.Create(req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.Response{
		Message: "success",
	})
}

func (controller *Controller) Delete(c echo.Context) error {

	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	err = controller.service.Ownership(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	err = controller.service.Remove(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.Response{
		Message: "success",
	})
}

func (controller *Controller) GetAll(c echo.Context) error {
	playlists, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.ResponseSuccess{
		Status: "success",
		Data:   playlists,
	})
}

func (controller *Controller) AddPlaylistMusic(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	createPlaylistMusicRequest := new(request.CreatePlaylistMusicRequset)
	if err := c.Bind(createPlaylistMusicRequest); err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	req := *createPlaylistMusicRequest.ToSpec(uint64(id))

	err = controller.service.AddPlaylistMusic(uint64(userId), req)
	fmt.Println(err)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.Response{
		Message: "success",
	})
}

func (controller *Controller) GetPlaylistMusicById(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	res, err := controller.service.GetPlaylistMusicById(uint64(userId), uint64(id))
	if err != nil {
		fmt.Println(err)
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.ResponseSuccess{
		Status: "success",
		Data:   res.Musics,
	})
}

func (controller *Controller) RemovePlaylistMusicById(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	createPlaylistMusicRequest := new(request.CreatePlaylistMusicRequset)
	if err := c.Bind(createPlaylistMusicRequest); err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}

	err = controller.service.RemovePlaylistMusicById(uint64(userId), createPlaylistMusicRequest.MusicID, uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.Response{
		Message: "success",
	})
}
