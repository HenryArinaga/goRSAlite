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

	recAppearance1 := canvas.NewRectangle(color.Black)
	recAppearance1.SetMinSize(fyne.NewSize(500, 50))

	Appearance := widget.NewLabel("Appearance")
	Appearance.TextStyle = fyne.TextStyle{Bold: true}
	appearanceHeading1 := container.New(layout.NewCenterLayout(), Appearance, layout.NewSpacer())
	appearanceBox := container.New(layout.NewVBoxLayout(), appearanceHeading1, recAppearance1, layout.NewSpacer())

	return appearanceBox
}
