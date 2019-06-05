package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"

	log "github.com/sirupsen/logrus"

	"drone/models"
)

const (
	PORT = 9999
)

func main() {
	sendMessages := []string{"ASCII", "PROGRAMMING", "PLUS"}
	current := 0
	var conn net.Conn = nil

	for {
		if conn == nil {
			conn, _ = net.Dial("tcp", fmt.Sprintf("localhost:%d", PORT))
			log.Infof("Access: %d\n", current)
		}

		// POST
		body, _ := json.Marshal(&models.Data{
			Time: current,
		})
		request, _ := http.NewRequest(
			"POST",
			"http://localhost:8888",
			bytes.NewReader(body),
		)
		err := request.Write(conn)
		if err != nil {
			log.Error(err)
		}

		// サーバから読み込む。タイムアウトはここでエラーになるのでリトライ
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			log.Error("timeout")
			conn = nil
			continue
		}

		// 結果を表示
		dump, _ := httputil.DumpResponse(response, true)
		fmt.Println(string(dump))

		// 全部送信完了していれば終了I
		current++
		if current == len(sendMessages) {
			break
		}
	}
}
