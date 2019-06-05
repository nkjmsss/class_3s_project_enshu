package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/nkjmsss/class_3s_project_enshu/models"
)

const (
	PORT = 1323
)

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		panic(err)
	}
	log.Infof("server is listening at port %d", PORT)

	for {
		conn, _ := listen.Accept()
		go func() {
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			// read request body
			body, err := ioutil.ReadAll(request.Body)
			defer request.Body.Close()
			if err != nil {
				log.Error(err)
			}

			// unmarshal body
			data := &models.Data{}
			if err := json.Unmarshal(body, data); err != nil {
				log.Error(err)
			}

			log.Info(data)

			response := http.Response{
				StatusCode: 200,
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
