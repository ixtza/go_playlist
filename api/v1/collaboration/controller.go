package collaboration

import (
	v1 "mini-clean/api/v1"
	"mini-clean/api/v1/collaboration/request"
	"mini-clean/api/v1/collaboration/response"
	collaborationUsecase "mini-clean/service/collaboration"
	playlistUsecase "mini-clean/service/playlist"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	serviceCollab   collaborationUsecase.Service
	servicePlaylist playlistUsecase.Service
}

func NewController(serviceCollab collaborationUsecase.Service, servicePlaylsit playlistUsecase.Service) *Controller {
	return &Controller{
		serviceCollab:   serviceCollab,
		servicePlaylist: servicePlaylsit,
	}
}

func (controller *Controller) Create(c echo.Context) error {
	createCollborationRequest := new(request.CreateCollborationRequest)
	if err := c.Bind(createCollborationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	req := *createCollborationRequest.ToSpec()

	data := strings.Split(c.Get("payload").(string), ":")

	ownerId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	err = controller.servicePlaylist.Ownership(uint64(ownerId), uint64(req.PlaylistID))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "error",
			Message: err.Error(),
		})

	}

	err = controller.serviceCollab.Create(req)

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: "User " + strconv.Itoa(int(createCollborationRequest.UserID)) + " added to collaboration",
	})
}

func (controller *Controller) Remove(c echo.Context) error {
	createCollborationRequest := new(request.CreateCollborationRequest)
	if err := c.Bind(createCollborationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	req := *createCollborationRequest.ToSpec()

	data := strings.Split(c.Get("payload").(string), ":")

	ownerId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	err = controller.servicePlaylist.Ownership(uint64(ownerId), uint64(req.PlaylistID))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	err = controller.serviceCollab.Remove(uint64(req.UserID), uint64(req.PlaylistID))

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	userId := strconv.Itoa(int(req.UserID))

	return c.JSON(v1.GetErrorStatus(err), response.Response{
		Status:  "success",
		Message: "User id's " + userId + " removed from collboration",
	})
}
