package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"zerotomastery.io/pixl/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected    bool
	Color       color.Color
	SwatchIndex int
	// Supplied code to change the swatch after mouse click
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(color color.Color) {
	s.Color = color
	s.Refresh()
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	s := &Swatch{
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}
	// Supplied all the state information to the base widget
	s.ExtendBaseWidget(s)
	return s
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  s,
	}
}
