package ui

import "github.com/keaysma/termui/v3"

func GetTerminalDimensions() TerminalDimensions {
	width, height := termui.TerminalDimensions()
	return TerminalDimensions{width, height}
}
