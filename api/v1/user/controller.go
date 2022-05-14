package user

import (
	v1 "mini-clean/api/v1"
	"mini-clean/api/v1/user/request"
	"mini-clean/api/v1/user/response"
	userUsecase "mini-clean/service/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service userUsecase.Service
}

func NewController(service userUsecase.Service) *Controller {
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
	user, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ResponseSuccess{
		Status: "success",
		Data:   user,
	})
}

func (controller *Controller) Modify(c echo.Context) error {

	params := c.Param("id")
	if params == "" {
		return c.JSON(http.StatusNotFound, response.Response{
			Status:  "fail",
			Message: "put user id in endpoint",
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	if data[0] != params {
		return c.JSON(http.StatusUnauthorized, response.Response{
			Status:  "fail",
			Message: "invalid username or password",
		})
	}

	createUserRequest := new(request.CreateUserRequest)
	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createUserRequest.ToSpec()
	_, err := controller.service.Modify(req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.Response{
		Status:  "success",
		Message: "user data updated",
	})
}

func (controller *Controller) GetAll(c echo.Context) error {

	users, err := controller.service.GetAll()

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ResponseSuccess{
		Status: "success",
		Data:   users,
	})
}

func (controller *Controller) Create(c echo.Context) error {
	createUserRequest := new(request.CreateUserRequest)
	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createUserRequest.ToSpec()

	err := controller.service.Create(req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.Response{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: "user created",
	})
}

// func (controller *Controller) Delete(c echo.Context) error {
// 	return errors.New("some error")
// }
