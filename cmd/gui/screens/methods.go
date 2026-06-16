package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MethodsScreen() fyne.CanvasObject {

	inside := container.New(layout.NewVBoxLayout(), widget.NewLabel("this is the methods screen"))

	return inside
}
