package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"zerotomastery.io/pixl/pxcanvas/brush"
)

// Scrolled Implement the fyne.Scrollable interface
func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

// MouseIn Implement the fyne.Hoverable interface
func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {

}

func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	// Click and drag to draw
	if x, y := pxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
		// Add cursor
		cursor := brush.Cursor(pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, ev, *x, *y)
		pxCanvas.renderer.SetCursor(cursor)
	} else {
		// Move the mouse outside canvas area, this will remove the cursor
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

func (pxCanvas *PxCanvas) MouseOut() {
}

// MouseUp Implement the fyne.Mouseable interface
func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {

}

func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}
