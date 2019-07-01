package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/nkjmsss/class_3s_project_enshu/middleware/gobot"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/handler"
)

const (
	PORT = 1323
)

func main() {
	// gobot.Init()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "status=${status}, error=${error}\n",
	}))

	e.GET("/", handler.HandleGet)
	e.POST("/", handler.HandlePost)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
