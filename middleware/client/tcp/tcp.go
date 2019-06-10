package tcp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
)

const (
	URI  = "controller"
	PORT = 1323
)

func SendTCP(data *models.Data) {
	conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", URI, PORT))

	// POST
	body, _ := json.Marshal(data)
	request, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("http://%s:%d", URI, PORT),
		bytes.NewReader(body),
	)
	err := request.Write(conn)
	if err != nil {
		log.Error(err)
	}
}
