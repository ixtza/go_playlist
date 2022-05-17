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

// GetByID godoc
// @Summary      Get music
// @Description  Retrive music data by id
// @Tags         musics
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "Music ID"
// @Success      200  {object}  response.MusicResponseSuccess
// @Failure      400  {object}  response.MusicResponse
// @Failure      403  {object}  response.MusicResponse
// @Failure      500  {object}  response.MusicResponse
// @Security ApiKeyAuth
// @Router       /musics/{id} [get]
func (controller *Controller) GetByID(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	music, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.MusicResponseSuccess{
		Status: "success",
		Data:   music,
	})
}

// Modify godoc
// @Summary      Modify music
// @Description  Edit music data by id
// @Tags         musics
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateMusicRequest true "Modify music json format" SchemaExample(request.CreateMusicRequest)
// @Success      200  {object}  response.MusicResponse
// @Failure      400  {object}  response.MusicResponse
// @Failure      403  {object}  response.MusicResponse
// @Failure      404  {object}  response.MusicResponse
// @Failure      500  {object}  response.MusicResponse
// @Security ApiKeyAuth
// @Router       /musics/{id} [put]
func (controller *Controller) Modify(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	createMusicRequest := new(request.CreateMusicRequest)
	if err := c.Bind(createMusicRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createMusicRequest.ToSpec()
	_, err = controller.service.Modify(uint64(id), req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.MusicResponse{
		Status:  "success",
		Message: "Music iD :" + params + " Updated",
	})
}

// Create godoc
// @Summary      Create music
// @Description  Insert new music
// @Tags         musics
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateMusicRequest true "Modify music json format" SchemaExample(request.CreateMusicRequest)
// @Success      201  {object}  response.MusicResponse
// @Failure      400  {object}  response.MusicResponse
// @Failure      403  {object}  response.MusicResponse
// @Failure      500  {object}  response.MusicResponse
// @Security ApiKeyAuth
// @Router       /musics [post]
func (controller *Controller) Create(c echo.Context) error {
	createMusicRequest := new(request.CreateMusicRequest)
	if err := c.Bind(createMusicRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createMusicRequest.ToSpec()

	id, err := controller.service.Create(req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MusicResponseSuccess{
		Status: "success",
		Data:   map[string]interface{}{"music_id": id},
	})
}

// Delete godoc
// @Summary      Delete music
// @Description  Remove music by id
// @Tags         musics
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "Music ID"
// @Success      200  {object}  response.MusicResponse
// @Failure      400  {object}  response.MusicResponse
// @Failure      403  {object}  response.MusicResponse
// @Failure      404  {object}  response.MusicResponse
// @Failure      500  {object}  response.MusicResponse
// @Security ApiKeyAuth
// @Router       /musics/{id} [delete]
func (controller *Controller) Delete(c echo.Context) error {

	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	err = controller.service.Remove(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.MusicResponse{
		Status:  "success",
		Message: "music deleted",
	})
}

// GetAll godoc
// @Summary      Get all musics
// @Description  Retrive all music data
// @Tags         musics
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.MusicResponse
// @Failure      403  {object}  response.MusicResponse
// @Failure      500  {object}  response.MusicResponse
// @Security ApiKeyAuth
// @Router       /musics [get]
func (controller *Controller) GetAll(c echo.Context) error {
	musics, err := controller.service.GetAll()
	if err != nil || len(musics) == 0 {
		return c.JSON(v1.GetErrorStatus(err), response.MusicResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.MusicResponseSuccess{
		Status: "success",
		Data:   musics,
	})
}
