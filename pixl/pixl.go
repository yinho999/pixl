package main

import (
	"fyne.io/fyne/v2/app"
	"image/color"
	"zerotomastery.io/pixl/apptype"
	"zerotomastery.io/pixl/swatch"
	"zerotomastery.io/pixl/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("Pixl")
	state := apptype.State{
		BrushColor:     color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		SwatchSelected: 0,
	}
	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()
}
