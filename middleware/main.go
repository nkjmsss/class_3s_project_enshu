package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/nkjmsss/class_3s_project_enshu/middleware/gobot"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/handler"
)

const (
	PORT = 1323
)

func init() {
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", 0777)
	}
}

func main() {
	// gobot.Init()

	e := echo.New()
	e.Use(middleware.Recover())

	logfile, err := os.OpenFile("./logs/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open ./logs/access.log:" + err.Error())
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Format: "status=${status}, error=${error}\n",
		Output: logfile,
	}))

	e.GET("/", handler.HandleGet)
	e.POST("/", handler.HandlePost)
	e.File("/", "public")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
