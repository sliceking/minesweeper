package main

import (
	"fmt"

	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("MineSweeper")
	w.SetContent(widget.NewTabContainer(
		widget.NewTabItem("game", makeGameBoard()),
	))

	w.ShowAndRun()
}

func makeGameBoard() *fyne.Container {
	gridWidth := 10
	gridArea := gridWidth * gridWidth
	cells := []fyne.CanvasObject{}

	for i := 0; i < gridArea; i++ {
		cells = append(cells, makeCell())
	}

	return fyne.NewContainerWithLayout(layout.NewGridLayout(gridWidth), cells...)
}

func makeCell() *widget.Button {
	rect := widget.NewButton("M", func() { fmt.Println("yay") })
	return rect
}
