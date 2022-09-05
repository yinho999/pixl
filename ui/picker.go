package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lusingander/colorpicker"
	"image/color"
)

func SetupColorPicker(app *AppInit) *fyne.Container {
	// Create a color picker
	picker := colorpicker.New(200, colorpicker.StyleValue)
	// Change selected swatch color when color picker value changes
	picker.SetOnChanged(func(c color.Color) {
		app.State.BrushColor = c
		app.Swatches[app.State.SwatchSelected].SetColor(c)
	})
	return container.NewVBox(picker)
}
