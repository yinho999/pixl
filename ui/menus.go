package ui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/png"
	"os"
	"strconv"
	"zerotomastery.io/pixl/util"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(file fyne.URIWriteCloser, e error) {
		if e != nil {
			dialog.ShowError(e, app.PixlWindow)
			return
		}
		if file == nil {
			return
		}
		err := png.Encode(file, app.PixlCanvas.PixelData)
		if err != nil {
			dialog.ShowError(err, app.PixlWindow)
			return
		}
		app.State.SetFilePath(file.URI().Path())

	}, app.PixlWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(file *os.File) {
				err := file.Close()
				if err != nil {
					dialog.ShowError(err, app.PixlWindow)
				}
			}
			file, err := os.Create(app.State.FilePath)
			defer tryClose(file)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
			err = png.Encode(file, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")
			}
			if width <= 0 {
				return errors.New("must be > 0 ")
			}
			return nil
		}
		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}
		dialog.ShowForm("New", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0
				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("Invalid width"), app.PixlWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}
				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("Invalid width"), app.PixlWindow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}
				app.PixlCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixlWindow)
	})
}

func BuildLoadMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Load...", func() {
		dialog.ShowFileOpen(func(file fyne.URIReadCloser, e error) {
			if e != nil {
				dialog.ShowError(e, app.PixlWindow)
				return
			}
			if file == nil {
				return
			}
			image, _, err := image.Decode(file)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
			app.PixlCanvas.LoadImage(image)
			app.State.SetFilePath(file.URI().Path())
			imgColors := util.GetImageColors(image)
			i := 0
			for color := range imgColors {
				if i == len(app.Swatches) {
					break
				}
				app.Swatches[i].SetColor(color)
				i++
			}

		}, app.PixlWindow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu("File", BuildNewMenu(app), BuildLoadMenu(app), BuildSaveMenu(app), BuildSaveAsMenu(app))
}

func SetupMenus(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixlWindow.SetMainMenu(mainMenu)
}
