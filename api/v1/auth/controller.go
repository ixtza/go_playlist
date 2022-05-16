package auth

import (
	v1 "mini-clean/api/v1"
	"mini-clean/api/v1/auth/request"
	"mini-clean/api/v1/auth/response"
	authService "mini-clean/service/auth"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service   authService.Service
	validator *validator.Validate
}

func NewController(service authService.Service) *Controller {
	return &Controller{
		service:   service,
		validator: validator.New(),
	}
}

// Auth godoc
// @Summary      User Authentication
// @Description  Authenticate user's info
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.AuthResponseSuccess
// @Failure      400  {object}  response.AuthResponse
// @Failure      401  {object}  response.AuthResponse
// @Failure      500  {object}  response.AuthResponse
// @Router       /auth [post]
func (controller *Controller) Auth(c echo.Context) error {
	authRequest := new(request.AuthRequest)
	if err := c.Bind(authRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.AuthResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	token, err := controller.service.Login(*authRequest.ToSpec())
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.AuthResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.AuthResponseSuccess{
		Status: "success",
		Data:   token,
	})
}
