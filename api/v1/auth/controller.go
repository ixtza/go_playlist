package auth

import (
	"mini-clean/api/v1/auth/request"
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

func (controller *Controller) Auth(c echo.Context) error {
	authRequest := new(request.AuthRequest)
	if err := c.Bind(authRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := controller.service.Login(*authRequest.ToSpec())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
