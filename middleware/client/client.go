package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/marcusolsson/tui-go"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/tcp"
	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
)

func main() {
	data := &models.Data{}

	ui, err := tui.New(tui.NewVBox(tui.NewLabel("controll your github.com/nkjmsss/class_3s_project_enshu with arrow key")))
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Left", func() {
		data.X--
		handleClick(data, ui)
	})
	ui.SetKeybinding("Right", func() {
		data.X++
		handleClick(data, ui)
	})
	ui.SetKeybinding("Up", func() {
		data.Y++
		handleClick(data, ui)
	})
	ui.SetKeybinding("Down", func() {
		data.Y--
		handleClick(data, ui)
	})
	ui.SetKeybinding("s", func() {
		data.Z++
		handleClick(data, ui)
	})
	ui.SetKeybinding("d", func() {
		data.Z--
		handleClick(data, ui)
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func updateUI(ui tui.UI, data *models.Data) {
	widget := tui.NewVBox(tui.NewLabel(data.String()))
	ui.SetWidget(widget)
}

func handleClick (data *models.Data, ui tui.UI) {
	data.Time = int(time.Now().UnixNano() / int64(time.Millisecond))
	updateUI(ui, data)
	if err := tcp.SendTCP(data, "localhost"); err != nil {
		log.Error(err)
	}
}
