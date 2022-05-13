package collaboration

import (
	"mini-clean/api/v1/collaboration/request"
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
		return c.JSON(http.StatusBadRequest, err)
	}

	req := *createCollborationRequest.ToSpec()

	data := strings.Split(c.Get("payload").(string), ":")

	ownerId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}
	err = controller.servicePlaylist.Ownership(uint64(ownerId), uint64(req.PlaylistID))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "User unauthorized")
	}

	err = controller.serviceCollab.Create(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func (controller *Controller) Remove(c echo.Context) error {
	createCollborationRequest := new(request.CreateCollborationRequest)
	if err := c.Bind(createCollborationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := *createCollborationRequest.ToSpec()

	data := strings.Split(c.Get("payload").(string), ":")

	ownerId, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}

	err = controller.servicePlaylist.Ownership(uint64(ownerId), uint64(req.PlaylistID))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "User unauthorized")
	}

	res, err := controller.serviceCollab.Remove(uint64(req.UserID), uint64(req.PlaylistID))

	if err != nil || res {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
