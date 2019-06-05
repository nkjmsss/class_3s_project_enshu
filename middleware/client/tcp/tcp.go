package tcp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
)

const (
	URI  = "localhost"
	PORT = 1323
)

func SendTCP(data *models.Data) {
	var conn net.Conn = nil
	maxtry := 3

	for i := 0; i < maxtry; i++ {
		if conn == nil {
			conn, _ = net.Dial("tcp", fmt.Sprintf("%s:%d", URI, PORT))
		}

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

		// サーバから読み込む。タイムアウトはここでエラーになるのでリトライ
		if _, err := http.ReadResponse(bufio.NewReader(conn), request); err != nil {
			log.Error("timeout")
			conn = nil
			continue
		}

		if i == maxtry-1 {
			log.Error("request failed")
		}

		// ここまでたどりつけばリクエスト成功
		break
	}
}
