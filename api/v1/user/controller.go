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

// GetByID godoc
// @Summary      Get user
// @Description  Retrive user data by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "User ID"
// @Success      200  {object}  response.UserResponseSuccess
// @Failure      400  {object}  response.UserResponse
// @Failure      403  {object}  response.UserResponse
// @Failure      404  {object}  response.UserResponse
// @Failure      500  {object}  response.UserResponse
// @Security ApiKeyAuth
// @Router       /users/{id} [get]
func (controller *Controller) GetByID(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	user, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.UserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.UserResponseSuccess{
		Status: "success",
		Data:   user,
	})
}

// Modify godoc
// @Summary      Modify user
// @Description  Retrive user data by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "User ID"
// @Param // @Param Payload body request.CreateUserRequest true "Playlist json format" SchemaExample(request.CreateUserRequest)
// @Success      200  {object}  response.UserResponse
// @Failure      400  {object}  response.UserResponse
// @Failure      401  {object}  response.UserResponse
// @Failure      403  {object}  response.UserResponse
// @Failure      404  {object}  response.UserResponse
// @Failure      500  {object}  response.UserResponse
// @Security ApiKeyAuth
// @Router       /users/{id} [put]
func (controller *Controller) Modify(c echo.Context) error {

	params := c.Param("id")
	if params == "" {
		return c.JSON(http.StatusNotFound, response.UserResponse{
			Status:  "fail",
			Message: "put user id in endpoint",
		})
	}

	data := strings.Split(c.Get("payload").(string), ":")

	if data[0] != params {
		return c.JSON(http.StatusUnauthorized, response.UserResponse{
			Status:  "fail",
			Message: "invalid username or password",
		})
	}

	createUserRequest := new(request.CreateUserRequest)
	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createUserRequest.ToSpec()
	_, err := controller.service.Modify(req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.UserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.UserResponse{
		Status:  "success",
		Message: "user data updated",
	})
}

func (controller *Controller) GetAll(c echo.Context) error {

	users, err := controller.service.GetAll()

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.UserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.UserResponseSuccess{
		Status: "success",
		Data:   users,
	})
}

// Create godoc
// @Summary      Create new user
// @Description  Retrive user data by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id   path  int  true  "User ID"
// @Param // @Param Payload body request.CreateUserRequest true "Playlist json format" SchemaExample(request.CreateUserRequest)
// @Success      200  {object}  response.UserResponseSuccess
// @Failure      400  {object}  response.UserResponse
// @Failure      500  {object}  response.UserResponse
// @Router       /signup [post]
func (controller *Controller) Create(c echo.Context) error {
	createUserRequest := new(request.CreateUserRequest)
	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	req := *createUserRequest.ToSpec()

	id, err := controller.service.Create(req)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.UserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.UserResponseSuccess{
		Status: "success",
		Data:   map[string]interface{}{"user_id": id},
	})
}

// func (controller *Controller) Delete(c echo.Context) error {
// 	return errors.New("some error")
// }
