package screens

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SettingsScreen() fyne.CanvasObject {

	lightModeCheck := widget.NewCheck("Light mode (coming soon)",
		func(turnOnLightMode bool) {

		})

	recAppearance1 := canvas.NewRectangle(color.Black)
	recAppearance1.SetMinSize(fyne.NewSize(500, 50))
	lightModeCheck.Disable()

	Appearance := widget.NewLabel("Appearance")
	Appearance.TextStyle = fyne.TextStyle{Bold: true}
	appearanceHeading1 := container.New(layout.NewCenterLayout(), Appearance, layout.NewSpacer())

	lightModeRow := container.New(layout.NewStackLayout(), recAppearance1, lightModeCheck)

	appearanceBox := container.New(layout.NewVBoxLayout(), appearanceHeading1, lightModeRow, layout.NewSpacer())

	centered1 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), appearanceBox, layout.NewSpacer())

	return centered1
}
