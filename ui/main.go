package ui

import (
	"log"

	"github.com/keaysma/termui/v3"
	"github.com/keaysma/termui/v3/widgets"
)

type TerminalDimensions struct {
	width  int
	height int
}

func Draw() {

}

func UI() {
	err := termui.Init()
	if err != nil {
		log.Fatalf("failed to initialize termui: %s", err)
	}
	defer termui.Close()
	td := GetTerminalDimensions()
	tabpane := widgets.NewTabPane()
	tabpane.SetRect(0, 0, int(float32(td.width)*0.75), 3)
	termui.Render(tabpane)
	for event := range termui.PollEvents() {
		//fmt.Printf("Event Type: %s, ID: %s\n", event.Type, event.ID)
		switch event.Type {
		case termui.ResizeEvent:

		case termui.KeyboardEvent:
			switch event.ID {
			case "q", "<C-c>":
				return
			case "<s>":
				tabpane.FocusLeft()
				termui.Clear()
				termui.Render(tabpane)
				//renderTab()
			case "<d>":
				tabpane.FocusRight()
				termui.Clear()
				termui.Render(tabpane)
				//renderTab()
			default:
			}
		default:
		}
	}
}
