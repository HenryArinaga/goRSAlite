package screens

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type State struct {
	SelectedMethod string
}

var state State

func MethodsScreen() fyne.CanvasObject {

	combo := widget.NewSelect(
		[]string{"Trial Division", "Square Root", "Fermat Factorization"},
		func(value string) {
			log.Println("Select set to", value)
			state.SelectedMethod = value

		})

	combo.SetSelected(state.SelectedMethod)

	inside := container.New(
		layout.NewVBoxLayout(),

		combo,
	)

	return inside
}
