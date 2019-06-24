package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	log "github.com/sirupsen/logrus"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/tcp"
)

const (
	PORT = 1323
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", handleGet)
	e.POST("/", handlePost)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}

func handleGet(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func handlePost(c echo.Context) error {
	d := new(models.Data)
	if err := c.Bind(d); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	d.Time = int(time.Now().UnixNano() / int64(time.Millisecond))

	if err := tcp.SendTCP(d, "controller"); err != nil {
		return err
	}

	log.Infof("{time: %d, x: %d, y: %d, z: %d, shape: %d}", d.Time, d.X, d.Y, d.Z, d.Shape)

	return c.JSON(http.StatusOK, d)
}
