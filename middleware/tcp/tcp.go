package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
)

const (
	DEFAULT_URI = "controller"
	PORT        = 1324
)

func SendTCP(data *models.SendData, uri string) error {
	URI := DEFAULT_URI
	if uri != "" {
		URI = uri
	}

	// d := net.Dialer{
	// 	Timeout: 10 * time.Millisecond,
	// }

	// conn, err := d.Dial("tcp", fmt.Sprintf("%s:%d", URI, PORT))
	// if err != nil {
	// 	return err
	// }
	// defer conn.Close()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", URI, PORT))
	if err != nil {
		fmt.Println("net resolve TCP Addr error ")
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("net dial tcp error ")
		os.Exit(1)
	}
	conn.SetKeepAlive(true)
	defer conn.Close()

	// POST
	_d := &struct {
		Time    int `json:"time"`
		X       int `json:"x"`
		Y       int `json:"y"`
		Z       int `json:"z"`
		Shape   int `json:"shape"`
		Command int `json:"command"`
	}{
		Time:    data.Time,
		X:       data.RightHand.X,
		Y:       data.RightHand.Y,
		Z:       data.RightHand.Z,
		Shape:   data.RightHand.Shape,
		Command: data.Command,
	}
	body, _ := json.Marshal(_d)
	// request, _ := http.NewRequest(
	// 	"POST",
	// 	fmt.Sprintf("http://%s:%d", URI, PORT),
	// 	bytes.NewReader(body),
	// )
	// if err := request.Write(conn); err != nil {
	// 	return err
	// }
	if _, err := conn.Write(body); err != nil {
		return err
	}

	// output command log
	logfile, err := os.OpenFile("./logs/command.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open ./logs/command.log:" + err.Error())
	}
	log.SetOutput(logfile)

	response, _ := bufio.NewReader(conn).ReadString('\n')
	if response != "" {
		log.Info(response)
	}

	return nil
}
