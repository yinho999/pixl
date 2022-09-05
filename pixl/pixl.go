package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"image/color"
	"zerotomastery.io/pixl/apptype"
	"zerotomastery.io/pixl/pxcanvas"
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

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea: fyne.NewSize(600, 600),
		// the offset is set to the top left corner of the drawing area
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		// Every one pixel in the canvas is 30 pixels in the on screen
		PxSize: 30,
	}
	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()
}
