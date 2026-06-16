package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GO RSA Lite")

	HomeButton := widget.NewButton("Home", func() {
		log.Println("tapped")
	})
	MethodsButton := widget.NewButton("Methods", func() {
		log.Println("Methods tapped")
	})
	SettingsButton := widget.NewButton("Settings", func() {
		log.Println("Settings tapped")
	})
	LogsButton := widget.NewButton("Logs", func() {
		log.Println("Logs tapped")
	})

	w.Resize(fyne.NewSize(1200, 800))

	leftSidebar := container.NewVBox(
		HomeButton,
		MethodsButton,
		SettingsButton,
		LogsButton,
	)

	centerTop := container.New(layout.NewCenterLayout(), canvas.NewText("Welcome to the RSA Factorization Tool!", color.White))

	rightTop := container.New(layout.NewHBoxLayout(), widget.NewLabel("Right Top"))

	topBar := container.NewBorder(nil, nil, nil, rightTop, centerTop)

	centerContent := container.New(layout.NewCenterLayout(), canvas.NewText("Welcome to the RSA Factorization Tool!", color.White))

	outsideBorder := container.NewBorder(topBar, nil, leftSidebar, nil, centerContent)

	w.SetContent(outsideBorder)

	w.ShowAndRun()

}
