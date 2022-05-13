package music

import (
	v1 "mini-clean/api/v1"
	"mini-clean/api/v1/music/request"
	"mini-clean/api/v1/music/response"
	musicUsecase "mini-clean/service/musixmatch"
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
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	music, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.ResponseSuccess{
		Status: "success",
		Data:   music,
	})
}

func (controller *Controller) Modify(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return err
	}
	createMusicRequest := new(request.CreateMusicRequest)
	if err := c.Bind(createMusicRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createMusicRequest.ToSpec()
	_, err = controller.service.Modify(uint64(id), req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.Response{
		Status:  "success",
		Message: "Music iD :" + params + " Updated",
	})
}

func (controller *Controller) Create(c echo.Context) error {
	createMusicRequest := new(request.CreateMusicRequest)
	if err := c.Bind(createMusicRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createMusicRequest.ToSpec()

	err := controller.service.Create(req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: "music created",
	})
}

func (controller *Controller) Delete(c echo.Context) error {

	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	err = controller.service.Remove(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.Response{
		Status:  "success",
		Message: "music deleted",
	})
}

func (controller *Controller) GetAll(c echo.Context) error {
	musics, err := controller.service.GetAll()
	if err != nil || len(musics) == 0 {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.ResponseSuccess{
		Status: "success",
		Data:   musics,
	})
}
