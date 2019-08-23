package main

import (
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

	// fmt.Println(mineMap)

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
	// gridArea := gridWidth * gridWidth
	cells := []fyne.CanvasObject{}

	for i := 0; i < 100; i++ {
		label := "clear"
		if mineMap[i] == true {
			label = "mine"
		}
		sq := newSquare(label)
		cells = append(cells, sq)
	}

	return fyne.NewContainerWithLayout(layout.NewGridLayout(gridWidth), cells...)
}

// Square represents a square on the gameboard, it is very close to a
// fyne button, so it's an alias with some extra functions attached
// type Square struct {
// 	*widget.Button
// }

// func newSquare(label string) *Square {
// 	s := widget.NewButton(label, func() { fmt.Println(label) })
// 	sq := Square{s}
// 	return &sq
// }

// TappedSecondary is trying to overwrite a method with the same name
// on widget.Button.
// func (s *Square) TappedSecondary(*fyne.PointEvent) {
// 	fmt.Println("nooo")
// }
