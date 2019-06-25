package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/gesture"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/tcp"
)

const (
	PORT = 1323
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "status=${status}, error=${error}\n",
	}))

	e.GET("/", handleGet)
	e.POST("/", handlePost)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}

func handleGet(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func handlePost(c echo.Context) error {
	// receive and bind data
	r := new(models.ReceiveData)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	r.Time = int(time.Now().UnixNano() / int64(time.Millisecond))

	// reshape data into SendData
	d := &models.SendData{
		Time:      r.Time,
		RightHand: r.Right,
	}

	// check if gesture is done
	gesture.Log(r)
	if gesture.DoTakeoff() {
		d.Command = models.TAKEOFF
		fmt.Println("take off!!!")
	}
	if gesture.DoLand() {
		d.Command = models.LAND
		fmt.Println("land!!!")
	}

	// log.Info("\n" + d.RightHand.String())

	if err := tcp.SendTCP(d, "controller"); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, d)
}
