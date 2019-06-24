// first of all, you need to install "mplayer"
package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
	"golang.org/x/net/websocket"
)

var conn *websocket.Conn
var ch = make(chan int)

func main() {
	go func() {
		http.Handle("/", websocket.Handler(wsHandler))
		if err := http.ListenAndServe(":3000", nil); err != nil {
			panic("ListenAndServe: " + err.Error())
		}
		receive(conn)
	}()

	drone := tello.NewDriver("8890")
	keys := keyboard.NewDriver()

	work := func() {
		/**
		  ドローンカメラアクセス
		*/
		mplayer := exec.Command("mplayer", "-fps", "60", "-")
		drone.SetVideoEncoderRate(0)
		mplayerIn, _ := mplayer.StdinPipe()
		if err := mplayer.Start(); err != nil {
			fmt.Println(err)
			return
		}

		drone.On(tello.ConnectedEvent, func(data interface{}) {
			fmt.Println("Connected")
			drone.StartVideo()
			drone.SetVideoEncoderRate(4)
			gobot.Every(100*time.Millisecond, func() {
				drone.StartVideo()
			})
		})

		drone.On(tello.VideoFrameEvent, func(data interface{}) {
			pkt := data.([]byte)
			if _, err := mplayerIn.Write(pkt); err != nil {
				fmt.Println(err)
			}
		})

		/**
		  ドローン制御
		*/
		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)

			if key.Key == keyboard.C {
				fmt.Println("Command Test")
			} else if key.Key == keyboard.T {
				fmt.Println("Take Off!")
				drone.TakeOff() //離陸
			} else if key.Key == keyboard.L {
				fmt.Println("Land")
				drone.Land()
			} else if key.Key == keyboard.W {
				fmt.Println("↑")
				drone.Forward(10) //前進
			} else if key.Key == keyboard.Z {
				fmt.Println("↓")
				drone.Backward(10) //後退
			} else if key.Key == keyboard.S {
				fmt.Println("→")
				drone.Right(10) //右へ
			} else if key.Key == keyboard.A {
				fmt.Println("←")
				drone.Left(10) //左へ
			} else if key.Key == keyboard.U {
				fmt.Println("Up")
				drone.Up(10) //上昇
			} else if key.Key == keyboard.D {
				fmt.Println("Down")
				drone.Down(10) //下降
			} else if key.Key == keyboard.F {
				fmt.Println("Front Flip")
				drone.FrontFlip() //フリップ
			} else if key.Key == keyboard.B {
				fmt.Println("Back Flip")
				drone.BackFlip() //バックフリップ
			} else {
				fmt.Println("keyboard event!", key, key.Char)
			}
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone, keys}, //droneとkeysの二つのアダプタ
		work,
	)

	robot.Start()
}

func wsHandler(ws *websocket.Conn) {
	conn = ws
	<-ch
}

func receive(c *websocket.Conn) {
	msg := make([]byte, 100)
	err := websocket.Message.Receive(c, msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Received: %s.\n", msg[:])
}
