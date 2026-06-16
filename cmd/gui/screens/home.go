package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen() fyne.CanvasObject {

	inside := container.New(layout.NewVBoxLayout(), widget.NewLabel("this is the home screen"))

	return inside
}
