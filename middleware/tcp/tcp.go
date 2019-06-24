package tcp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
)

const (
	DEFAULT_URI = "controller"
	PORT        = 1323
)

func SendTCP(data *models.SendData, uri string) error {
	URI := DEFAULT_URI
	if uri != "" {
		URI = uri
	}

	d := net.Dialer{
		Timeout: 10 * time.Millisecond,
	}

	conn, err := d.Dial("tcp", fmt.Sprintf("%s:%d", URI, PORT))
	if err != nil {
		return err
	}
	defer conn.Close()

	// POST
	_d := &struct {
		Time  int `json:"time"`
		X     int `json:"x"`
		Y     int `json:"y"`
		Z     int `json:"z"`
		Shape int `json:"shape"`
	}{
		Time:  data.Time,
		X:     data.RightHand.X,
		Y:     data.RightHand.Y,
		Z:     data.RightHand.Z,
		Shape: data.RightHand.Shape,
	}
	body, _ := json.Marshal(_d)
	request, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("http://%s:%d", URI, PORT),
		bytes.NewReader(body),
	)
	if err := request.Write(conn); err != nil {
		return err
	}

	return nil
}
