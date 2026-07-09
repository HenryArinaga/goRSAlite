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

	mainRectangle := canvas.NewRectangle(color.Black)
	mainRectangle.SetMinSize(fyne.NewSize(500, 50))
	trialDivison := widget.NewLabel("Trial Divison")
	sqrt := widget.NewLabel("Square Root")
	fermat := widget.NewLabel("Fermat Method")

	content := container.New(layout.NewStackLayout(), mainRectangle, trialDivison)
	content2 := container.New(layout.NewStackLayout(), mainRectangle, sqrt)
	content3 := container.New(layout.NewStackLayout(), mainRectangle, fermat)

	centered1 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), content, layout.NewSpacer())
	centered2 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), content2, layout.NewSpacer())
	centered3 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), content3, layout.NewSpacer())

	cards := container.New(layout.NewVBoxLayout(), centered1, centered2, centered3)

	return cards
}
