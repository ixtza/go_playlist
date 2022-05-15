package main

import (
	"fmt"
	"log"
	"mini-clean/api"
	"mini-clean/app/modules"
	"mini-clean/config"
	"mini-clean/util"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title          go_playlist API
// @version        1.0
// @description    API yang memberikan layanan kepada pengguna untuk membuat playlist dan berkolaborasi
// @contact.name   API Support
// @contact.email  skiyejpjm@mozmail.com
// @Basepath       /v1
func main() {
	config := config.GetConfig()

	dbCon := util.NewConnectionDatabase(config)
	defer dbCon.CloseConnection()

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()

	handlSwag := echoSwagger.WrapHandler

	e.GET("swagger/*", handlSwag)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "go_playlist API")
	})

	api.RegistrationPath(e, controllers)

	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	<-quit
}
