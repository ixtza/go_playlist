package user

import (
	"errors"
	"mini-clean/api/v1/user/request"
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
		return err
	}
	user, err := controller.service.GetById(uint64(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *Controller) Modify(c echo.Context) error {
	createUserRequest := new(request.CreateUserRequest)
	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := strings.Split(c.Get("payload").(string), ":")

	if data[1] != createUserRequest.Email {
		return errors.New("invalid username or password")
	}

	req := *createUserRequest.ToSpec()
	_, err := controller.service.Modify(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) GetAll(c echo.Context) error {

	users, err := controller.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (controller *Controller) Create(c echo.Context) error {
	createUserRequest := new(request.CreateUserRequest)
	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := *createUserRequest.ToSpec()

	err := controller.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func (controller *Controller) Delete(c echo.Context) error {
	return errors.New("some error")
}
