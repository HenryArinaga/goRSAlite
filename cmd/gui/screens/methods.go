package screens

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MethodsScreen() fyne.CanvasObject {

	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	})
	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})

	inside := container.New(
		layout.NewVBoxLayout(),
		check,
		radio,
		combo,
	)

	return inside
}
