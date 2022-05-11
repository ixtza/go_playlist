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
)

func main() {
	config := config.GetConfig()

	dbCon := util.NewConnectionDatabase(config)
	defer dbCon.CloseConnection()

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Golang clean architecture")
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
