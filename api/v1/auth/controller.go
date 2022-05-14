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

func (controller *Controller) Auth(c echo.Context) error {
	authRequest := new(request.AuthRequest)
	if err := c.Bind(authRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	token, err := controller.service.Login(*authRequest.ToSpec())
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	// TODO: ganti status jadi accepted, data token tidak usah ditaruh kedalam interface
	return c.JSON(v1.GetErrorStatus(err), response.ResponseSuccess{
		Status: "fail",
		Data:   token,
	})
}
