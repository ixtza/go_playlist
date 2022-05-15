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

// Create godoc
// @Summary      Create collaboration
// @Description  Add user to collaboration
// @Tags         collaborations
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateCollborationRequest true "Add collaboration request json" SchemaExample(request.CreateCollborationRequest)
// @Success      201  {object}  response.CollaborationResponse
// @Failure      400  {object}  response.CollaborationResponse
// @Failure      401  {object}  response.CollaborationResponse
// @Failure      403  {object}  response.CollaborationResponse
// @Failure      500  {object}  response.CollaborationResponse
// @Security ApiKeyAuth
// @Router       /collaborations [post]
func (controller *Controller) Create(c echo.Context) error {
	createCollborationRequest := new(request.CreateCollborationRequest)
	if err := c.Bind(createCollborationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.CollaborationResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	req := *createCollborationRequest.ToSpec()

	data := strings.Split(c.Get("payload").(string), ":")

	ownerId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CollaborationResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	err = controller.servicePlaylist.Ownership(uint64(ownerId), uint64(req.PlaylistID))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.CollaborationResponse{
			Status:  "error",
			Message: err.Error(),
		})

	}

	err = controller.serviceCollab.Create(req)

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.CollaborationResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.CollaborationResponse{
		Status:  "success",
		Message: "User " + strconv.Itoa(int(createCollborationRequest.UserID)) + " added to collaboration",
	})
}

// Remove godoc
// @Summary      Remove collaboree
// @Description  Remove user from collaboree
// @Tags         collaborations
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateCollborationRequest true "Remove collaboration request json" SchemaExample(request.CreateCollborationRequest)
// @Success      200  {object}  response.CollaborationResponse
// @Failure      400  {object}  response.CollaborationResponse
// @Failure      401  {object}  response.CollaborationResponse
// @Failure      403  {object}  response.CollaborationResponse
// @Failure      404  {object}  response.CollaborationResponse
// @Failure      500  {object}  response.CollaborationResponse
// @Security ApiKeyAuth
// @Router       /collaborations [delete]
func (controller *Controller) Remove(c echo.Context) error {
	createCollborationRequest := new(request.CreateCollborationRequest)
	if err := c.Bind(createCollborationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.CollaborationResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	req := *createCollborationRequest.ToSpec()

	data := strings.Split(c.Get("payload").(string), ":")

	ownerId, err := strconv.Atoi(data[0])

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CollaborationResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	err = controller.servicePlaylist.Ownership(uint64(ownerId), uint64(req.PlaylistID))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.CollaborationResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	err = controller.serviceCollab.Remove(uint64(req.UserID), uint64(req.PlaylistID))

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.CollaborationResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	userId := strconv.Itoa(int(req.UserID))

	return c.JSON(v1.GetErrorStatus(err), response.CollaborationResponse{
		Status:  "success",
		Message: "User id's " + userId + " removed from collboration",
	})
}
