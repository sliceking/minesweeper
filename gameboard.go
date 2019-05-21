package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

type gameboardRenderer struct {
	icon  *canvas.Image
	label *canvas.Text

	objects   []fyne.CanvasObject
	gameboard *Gameboard
}

// MinSize calculates the minimum size of the gameboard
func (g *gameboardRenderer) MinSize() fyne.Size {
	min := fyne.NewSize(30, 30)
	return min
}

// Layout the components of the gameboard
func (g *gameboardRenderer) Layout(size fyne.Size) {
	return
}

// ApplyTheme is called when the gameboard needs to update its look
func (g *gameboardRenderer) ApplyTheme() {
	return
}

func (g *gameboardRenderer) BackgroundColor() color.Color {
	return theme.PrimaryColor()
}

func (g *gameboardRenderer) Refresh() {
	g.Layout(g.gameboard.Size())
	canvas.Refresh(g.gameboard)
}

func (g *gameboardRenderer) Objects() []fyne.CanvasObject {
	return g.objects
}

func (g *gameboardRenderer) Destroy() {

}

// Gameboard struct is sort of a copy of the button widget right now
type Gameboard struct {
	widget.baseWidget
	Text  string
	style ButtonStyle
	Icon  fyne.Resource

	OnTapped func() `json:"-"`
}
