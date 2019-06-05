package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/marcusolsson/tui-go"

	"github.com/nkjmsss/class_3s_project_enshu/client/tcp"
	"github.com/nkjmsss/class_3s_project_enshu/models"
)

func main() {
	data := &models.Data{}

	ui, err := tui.New(tui.NewVBox(tui.NewLabel("controll your github.com/nkjmsss/class_3s_project_enshu with arrow key")))
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Left", func() {
		data.X -= 1
		data.Time = int(time.Now().UnixNano() / int64(time.Millisecond))
		tcp.SendTCP(data)
		updateUI(ui, data)
	})
	ui.SetKeybinding("Right", func() {
		data.X += 1
		data.Time = int(time.Now().UnixNano() / int64(time.Millisecond))
		tcp.SendTCP(data)
		updateUI(ui, data)
	})
	ui.SetKeybinding("Up", func() {
		data.Y += 1
		data.Time = int(time.Now().UnixNano() / int64(time.Millisecond))
		tcp.SendTCP(data)
		updateUI(ui, data)
	})
	ui.SetKeybinding("Down", func() {
		data.Y -= 1
		data.Time = int(time.Now().UnixNano() / int64(time.Millisecond))
		tcp.SendTCP(data)
		updateUI(ui, data)
	})
	ui.SetKeybinding("s", func() {
		data.Z += 1
		data.Time = int(time.Now().UnixNano() / int64(time.Millisecond))
		tcp.SendTCP(data)
		updateUI(ui, data)
	})
	ui.SetKeybinding("d", func() {
		data.Z -= 1
		data.Time = int(time.Now().UnixNano() / int64(time.Millisecond))
		tcp.SendTCP(data)
		updateUI(ui, data)
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func updateUI(ui tui.UI, data *models.Data) {
	widget := tui.NewVBox(tui.NewLabel(data.String()))
	ui.SetWidget(widget)
}
