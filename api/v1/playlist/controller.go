package playlist

import (
	"fmt"
	"mini-clean/api/v1/playlist/request"
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
		return err
	}

	playlist, err := controller.service.GetById(uint64(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, playlist)
}

func (controller *Controller) Modify(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return err
	}
	createPlaylistRequest := new(request.CreatePlaylistRequest)
	if err := c.Bind(createPlaylistRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = controller.service.Ownership(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "User unauthorized")
	}

	req := *createPlaylistRequest.ToSpec(uint64(userId))
	_, err = controller.service.Modify(uint64(id), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) Create(c echo.Context) error {

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}

	createPlaylistRequest := new(request.CreatePlaylistRequest)
	if err := c.Bind(createPlaylistRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	req := *createPlaylistRequest.ToSpec(uint64(userId))
	err = controller.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func (controller *Controller) Delete(c echo.Context) error {

	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return err
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}

	err = controller.service.Ownership(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "User unauthorized")
	}

	err = controller.service.Remove(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) GetAll(c echo.Context) error {
	playlists, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, playlists)
}

func (controller *Controller) AddPlaylistMusic(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return err
	}

	data := strings.Split(c.Get("payload").(string), ":")
	fmt.Println(data)

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}

	createPlaylistMusicRequest := new(request.CreatePlaylistMusicRequset)
	if err := c.Bind(createPlaylistMusicRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	req := *createPlaylistMusicRequest.ToSpec(uint64(id))

	err = controller.service.AddPlaylistMusic(uint64(userId), req)
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) GetPlaylistMusicById(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println(err)
		return err
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}
	res, err := controller.service.GetPlaylistMusicById(uint64(userId), uint64(id))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) RemovePlaylistMusicById(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return err
	}

	data := strings.Split(c.Get("payload").(string), ":")

	userId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}
	err = controller.service.Remove(uint64(userId), uint64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}
