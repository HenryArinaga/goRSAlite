package screens

import (
	"fmt"
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

	content := container.NewVBox()
	for _, entry := range appController.History {
		text := fmt.Sprintf(
			"Time: %s\nType: %s\nInput: %s\nResult: %s\nMethod: %s\nDuration: %s\n",
			entry.Time.Format("15:04:05"),
			entry.Kind,
			entry.Input,
			entry.Result,
			entry.Method,
			entry.Duration,
		)
		label := widget.NewLabel(text)
		label.Wrapping = fyne.TextWrapWord
		content.Add(label)

	}

	scroll := container.NewVScroll(content)
	scroll.SetMinSize(fyne.NewSize(500, 500))

	centered := container.New(layout.NewVBoxLayout(), display, scroll, layout.NewSpacer())

	return centered
}
