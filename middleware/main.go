package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/tcp"
)

const (
	PORT = 1323
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.POST("/", handlePost)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}

func handlePost(c echo.Context) error {
	d := new(models.Data)
	if err := c.Bind(d); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := tcp.SendTCP(d, "controller"); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, d)
}
