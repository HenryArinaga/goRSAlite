package screens

import (
	"goRSAlite/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func LogsScreen(appController *controller.Controller) fyne.CanvasObject {

	inside := container.New(layout.NewVBoxLayout(), widget.NewLabel("this is the logs screen"))

	return inside
}
