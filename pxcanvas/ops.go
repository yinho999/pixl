package pxcanvas

import "fyne.io/fyne/v2"

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize -= 1
		}
	// if the direction is 0, do nothing
	default:
		pxCanvas.PxSize = 10
	}
}

// Pan Calculate the new pan position based on the previousCoordinates and the currentCoordinates
func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff

	pxCanvas.Refresh()
}
