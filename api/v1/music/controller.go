package music

import (
	"mini-clean/api/v1/music/request"
	musicUsecase "mini-clean/service/music"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service musicUsecase.Service
}

func NewController(service musicUsecase.Service) *Controller {
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
	music, err := controller.service.GetById(uint64(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, music)
}

func (controller *Controller) Modify(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return err
	}
	createMusicRequest := new(request.CreateMusicRequest)
	if err := c.Bind(createMusicRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// data := strings.Split(c.Get("payload").(string), ":")

	// userId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}

	req := *createMusicRequest.ToSpec()
	_, err = controller.service.Modify(uint64(id), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) Create(c echo.Context) error {
	createMusicRequest := new(request.CreateMusicRequest)
	if err := c.Bind(createMusicRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := *createMusicRequest.ToSpec()

	err := controller.service.Create(req)
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

	res, err := controller.service.Remove(uint64(id))
	if err != nil || res {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) GetAll(c echo.Context) error {
	musics, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, musics)
}
