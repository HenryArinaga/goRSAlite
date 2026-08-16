package screens

import (
	"fmt"
	"goRSAlite/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func LogsScreen(appController *controller.Controller) (fyne.CanvasObject, func()) {

	logsHeading1 := widget.NewLabel("Operaration History")
	logsHeading1.TextStyle = fyne.TextStyle{Bold: true}
	logsCentered := container.New(layout.NewCenterLayout(), logsHeading1)
	display := container.New(layout.NewVBoxLayout(), logsCentered)

	content := container.NewVBox()
	scroll := container.NewVScroll(content)
	refreshLogs := func() {
		content.RemoveAll()
		for _, entry := range appController.History {
			text := fmt.Sprintf(
				"Time: %s\nType: %s\nInput: %s\nResult: %s\nMethod: %s\nDuration: %s\nSieve: %t\nSIMD: %t\nExtended GCD: %t\n",
				entry.Time.Format("15:04:05"),
				entry.Kind,
				entry.Input,
				entry.Result,
				entry.Method,
				entry.Duration,
				entry.Sieve,
				entry.SIMD,
				entry.ExtendedGcd,
			)
			label := widget.NewLabel(text)
			label.Wrapping = fyne.TextWrapWord
			content.Add(label)
		}
		content.Refresh()
		scroll.ScrollToOffset(appController.LogsScrollOffset)
	}

	scroll.SetMinSize(fyne.NewSize(500, 500))

	scroll.OnScrolled = func(pos fyne.Position) {
		appController.LogsScrollOffset = pos
		fmt.Println("saved offset:", pos)

	}
	clearButton := container.New(layout.NewHBoxLayout(),
		widget.NewButton("Clear", func() {
			appController.History = nil
			refreshLogs()
		}),
	)
	topRow := container.NewBorder(nil, nil, nil, clearButton)

	centered := container.New(layout.NewVBoxLayout(), display, scroll, topRow, layout.NewSpacer())

	return centered, refreshLogs
}
