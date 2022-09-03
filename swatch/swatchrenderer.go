package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// SwatchRenderer Implement the WidgetRenderer interface
type SwatchRenderer struct {
	// CanvasObject
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

// MinSize Swatch minsize same size as the square
func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

// Layout Where the layout that the swatch will be drawn
// Resize the square to the parameter size
func (renderer *SwatchRenderer) Layout(size fyne.Size) {
	renderer.objects[0].Resize(size)
}

// Refresh internal state of the swatch
func (renderer *SwatchRenderer) Refresh() {
	renderer.Layout(fyne.NewSize(20, 20))
	renderer.square.FillColor = renderer.parent.Color
	// Indicate to the user that the swatch is selected
	if renderer.parent.Selected {
		// CanvasObject interface does not have access to StrokeWidth and StrokeColor properties
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{255, 255, 255, 255}
		// reassign the square to the objects array
		renderer.objects[0] = &renderer.square
	} else {
		renderer.square.StrokeWidth = 0
		renderer.objects[0] = &renderer.square
	}
	// Refresh the canvas objects, redraw the widget
	canvas.Refresh(renderer.parent)
}

// Objects return the objects that the swatch will draw
func (renderer *SwatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

// Destroy the swatch
func (renderer *SwatchRenderer) Destroy() {
}
