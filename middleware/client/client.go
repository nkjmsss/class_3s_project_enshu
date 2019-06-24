package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/marcusolsson/tui-go"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
)

func main() {
	data := &models.SendData{}

	ui, err := tui.New(tui.NewVBox(tui.NewLabel("controll your github.com/nkjmsss/class_3s_project_enshu with arrow key")))
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Left", func() {
		data.RightHand.X--
		handleClick(data, ui)
	})
	ui.SetKeybinding("Right", func() {
		data.RightHand.X++
		handleClick(data, ui)
	})
	ui.SetKeybinding("Up", func() {
		data.RightHand.Y++
		handleClick(data, ui)
	})
	ui.SetKeybinding("Down", func() {
		data.RightHand.Y--
		handleClick(data, ui)
	})
	ui.SetKeybinding("k", func() {
		data.RightHand.Z++
		handleClick(data, ui)
	})
	ui.SetKeybinding("j", func() {
		data.RightHand.Z--
		handleClick(data, ui)
	})

	ui.SetKeybinding("0", func() {
		data.RightHand.Shape = 0
		handleClick(data, ui)
	})
	ui.SetKeybinding("1", func() {
		data.RightHand.Shape = 1
		handleClick(data, ui)
	})
	ui.SetKeybinding("2", func() {
		data.RightHand.Shape = 2
		handleClick(data, ui)
	})
	ui.SetKeybinding("3", func() {
		data.RightHand.Shape = 3
		handleClick(data, ui)
	})
	ui.SetKeybinding("4", func() {
		data.RightHand.Shape = 4
		handleClick(data, ui)
	})
	ui.SetKeybinding("5", func() {
		data.RightHand.Shape = 5
		handleClick(data, ui)
	})
	ui.SetKeybinding("6", func() {
		data.RightHand.Shape = 6
		handleClick(data, ui)
	})
	ui.SetKeybinding("7", func() {
		data.RightHand.Shape = 7
		handleClick(data, ui)
	})
	ui.SetKeybinding("8", func() {
		data.RightHand.Shape = 8
		handleClick(data, ui)
	})
	ui.SetKeybinding("9", func() {
		data.RightHand.Shape = 9
		handleClick(data, ui)
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func updateUI(ui tui.UI, data *models.SendData) {
	widget := tui.NewVBox(tui.NewLabel(data.String()))
	ui.SetWidget(widget)
}

func handleClick(data *models.SendData, ui tui.UI) {
	// data.RightHand.Time = int(time.Now().UnixNano() / int64(time.Millisecond))
	updateUI(ui, data)
	if err := sendPost(data); err != nil {
		log.Error(err)
	}
}

func sendPost(data *models.SendData) error {
	jsonstr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"http://localhost:1323",
		bytes.NewBuffer(jsonstr),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
