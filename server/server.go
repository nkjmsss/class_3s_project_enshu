package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"drone/models"
)

const (
	PORT = 9999
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
			log.Infof("remote address: %v", conn.RemoteAddr())

			for {
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// タイムアウトもしくはソケットクローズ時は終了
					// それ以外はエラーにする
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						log.Error("timeout")
						break
					} else if err == io.EOF {
						break
					}
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
			}
			conn.Close()
		}()
	}
}
