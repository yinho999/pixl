package apptype

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
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

// Brushable will be implemented on PxCanvas
type Brushable interface {
	// Set color on x,y position
	SetColor(c color.Color, x, y int)
	// Convert Mouse event to x,y position
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}
