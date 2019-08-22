package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var mineMap [100]bool

func main() {
	// Generate a mine map
	mineMap = generateMineMap()

	fmt.Println(mineMap)

	// Start up a fyne application
	app := app.New()

	// Create the new window to work inside
	w := app.NewWindow("MineSweeper")

	// Add the menu and container to the window

	// create menu

	// create Container

	// add cells
	//

	//
	w.SetContent(widget.NewTabContainer(
		widget.NewTabItem("game", makeGameBoard()),
	))

	w.ShowAndRun()
}

// Creates a 100
func generateMineMap() [100]bool {
	var sb [100]bool

	// Create a random number generator using the time as a seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range sb {
		// Generate a random number between 1 and 100
		// and if it is greater than 60, mark it as a mine
		mine := r.Intn(100)

		if mine > 60 {
			sb[i] = true
			continue
		}

		sb[i] = false
	}

	return sb
}

func makeGameBoard() *fyne.Container {
	gridWidth := 10
	gridArea := gridWidth * gridWidth
	cells := []fyne.CanvasObject{}

	for i := 0; i < gridArea; i++ {
		cells = append(cells, newSquare())
	}

	return fyne.NewContainerWithLayout(layout.NewGridLayout(gridWidth), cells...)
}

func makeCell() *widget.Button {
	rect := widget.NewButton("M", func() { fmt.Println("yay") })
	return rect
}

// Square represents a square on the gameboard, it is very close to a
// fyne button, so it's an alias with some extra functions attached
type Square struct {
	*widget.Button
}

func newSquare() *Square {
	s := widget.NewButton("M", func() { fmt.Println("yay") })
	sq := Square{s}
	return &sq
}

// This is trying to overwrite the tappedsecondary method of the
// button widget.
func (s *Square) TappedSecondary(*fyne.PointEvent) {
	fmt.Println("nooo")
}
