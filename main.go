package main

import (
	"artifacts/game"

	"github.com/keaysma/termui/v3"
)

func main() {
	go game.GameLoop()
	for event := range termui.PollEvents() {
		switch event.Type {
		case termui.KeyboardEvent:
			switch event.ID {
			case "q", "<C-c>":
				return
			}
		default:
		}
	}
}
