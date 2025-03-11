package ui

import (
	"fyne.io/fyne/v2"
	"github.com/MSMITH71910/pixl/apptype"
	"github.com/MSMITH71910/pixl/pxcanvas"
	"github.com/MSMITH71910/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
