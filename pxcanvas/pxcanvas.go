package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"zerotomastery.io/pixl/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	PixelData  image.Image
	renderer   *PxCanvasRenderer
	mouseState PxCanvasMouseState
	appState   *apptype.State
	// to reload image that stored in PixelData
	reloadImage bool
}

// Bounds returns the rectangle bounding box inside the canvas
func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	// top left corner of the canvas
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	// bottom right corner of the canvas
	x1 := int(pxCanvas.PxCols*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(pxCanvas.PxRows*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(position fyne.Position, bounds image.Rectangle) bool {
	return position.X >= float32(bounds.Min.X) && position.X < float32(bounds.Max.X) && position.Y >= float32(bounds.Min.Y) && position.Y < float32(bounds.Max.Y)
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	// create a new blank image
	// NRGBA allow to set the alpha channel independently of other colors
	// RGBA multiply the color by the alpha channel
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	// fill the image with color
	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	img := NewBlankImage(config.PxCols, config.PxRows, color.NRGBA{R: 128, G: 128, B: 128, A: 255})
	// create a new canvas
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		PixelData:      img,
		appState:       state,
	}
	// set the canvas size
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}
	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}
	pxCanvas.renderer = renderer
	return renderer
}

// TryPan Point event is the mouse location and mouse event is button pressed
func (pxCanvas *PxCanvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	// if the previous coordinate is not nil and the mouse scroll button is pressed
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		pxCanvas.Pan(*previousCoord, ev.PointEvent)
	}
}

// SetColor set the color of the pixel at the given position
func (pxCanvas *PxCanvas) SetColor(c color.Color, x, y int) {
	// Does canvas has an user input image with NRGBA format?
	if nrgba, ok := pxCanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}
	// Does canvas has an user input image with RGBA format?
	if rgba, ok := pxCanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}

	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := pxCanvas.Bounds()
	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}
	pxSize := float32(pxCanvas.PxSize)
	xOffset := pxCanvas.CanvasOffset.X
	yOffset := pxCanvas.CanvasOffset.Y

	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}

func (pxCanvas *PxCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()
	pxCanvas.PxCanvasConfig.PxCols = dimensions.Dx()
	pxCanvas.PxCanvasConfig.PxRows = dimensions.Dy()

	pxCanvas.PixelData = img
	pxCanvas.reloadImage = true

	pxCanvas.Refresh()
}
func (pxCanvas *PxCanvas) NewDrawing(cols, rows int) {
	pxCanvas.appState.SetFilePath("")
	pxCanvas.PxCols = cols
	pxCanvas.PxRows = rows

	pixelData := NewBlankImage(cols, rows, color.NRGBA{R: 128, G: 128, B: 128, A: 255})

	pxCanvas.LoadImage(pixelData)
}
