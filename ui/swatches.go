package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"image/color"
	"zerotomastery.io/pixl/swatch"
)

// BuildSwatches
// containers contains the list of layouts
func BuildSwatches(app *AppInit) *fyne.Container {
	// Buffer with a size of 64, with all initialized to 0
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	// Capacity of app.Swatches
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		s := swatch.NewSwatch(app.State, initialColor, i, func(s *swatch.Swatch) {
			// Remove any selection in app.Swatches
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				app.Swatches[j].Refresh()
			}
			// Highlight the selected swatch
			app.State.SwatchSelected = s.SwatchIndex
			// Paint with the color
			app.State.BrushColor = s.Color
		})
		// Select the zero index swatch to be selected
		if i == 0 {
			s.Selected = true
			app.State.SwatchSelected = 0
			s.Refresh()
		}
		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}
	// Create a container with the swatches using GridWrapLayout
	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
