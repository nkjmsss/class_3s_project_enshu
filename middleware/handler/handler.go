package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/history"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/tcp"
)

func HandleGet(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func HandlePost(c echo.Context) error {
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

	history.Log(r)

	// check if gesture is done
	if history.DoTakeoff() {
		d.Command = models.TAKEOFF
		// fmt.Println("take off!!!")
	}
	if history.DoLand() {
		d.Command = models.LAND
		// fmt.Println("land!!!")
	}

	// log.Info("\n" + d.RightHand.String())

	if send, shape := history.DoSend(); send {
		d.RightHand.Shape = shape

		if err := sendTCP(d); err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, d)
}

func sendTCP(d *models.SendData) error {
	return tcp.SendTCP(d, "controller")
}
