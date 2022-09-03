package apptype

import (
	"fyne.io/fyne/v2"
	"image/color"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	// PxSize = 5 is five pixel on screen for each pixel in canvas
	PxSize int
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

func (s *State) SetFilePath(path string) {
	s.FilePath = path
}
