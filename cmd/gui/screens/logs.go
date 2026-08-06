package screens

import (
	"goRSAlite/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func LogsScreen(appController *controller.Controller) fyne.CanvasObject {

	logsHeading1 := widget.NewLabel("Operaration Histroy")
	logsHeading1.TextStyle = fyne.TextStyle{Bold: true}
	logsCentered := container.New(layout.NewCenterLayout(), logsHeading1)
	display := container.New(layout.NewVBoxLayout(), logsCentered)

	return display
}
