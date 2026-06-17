package main

import (
	"goRSAlite/cmd/gui/screens"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GO RSA Lite")

	centerTop := container.New(layout.NewCenterLayout(), canvas.NewText("Welcome to the RSA Factorization Tool!", color.White))

	rightTop := container.New(layout.NewHBoxLayout(), widget.NewLabel("Right Top"))

	topBar := container.NewBorder(nil, nil, nil, rightTop, centerTop)

	centerContent := container.New(layout.NewStackLayout(), screens.HomeScreen())

	HomeButton := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		log.Println("tapped")
		centerContent.Objects = []fyne.CanvasObject{screens.HomeScreen()}
		centerContent.Refresh()
	})
	MethodsButton := widget.NewButton("Methods", func() {
		log.Println("Methods tapped")
		centerContent.Objects = []fyne.CanvasObject{screens.MethodsScreen()}
		centerContent.Refresh()
	})
	SettingsButton := widget.NewButton("Settings", func() {
		log.Println("Settings tapped")
		centerContent.Objects = []fyne.CanvasObject{screens.SettingsScreen()}
		centerContent.Refresh()
	})
	LogsButton := widget.NewButton("Logs", func() {
		log.Println("Logs tapped")
		centerContent.Objects = []fyne.CanvasObject{screens.LogsScreen()}
		centerContent.Refresh()
	})

	w.Resize(fyne.NewSize(1200, 800))

	leftSidebar := container.NewVBox(
		HomeButton,
		MethodsButton,
		SettingsButton,
		LogsButton,
	)

	outsideBorder := container.NewBorder(topBar, nil, leftSidebar, nil, centerContent)

	w.SetContent(outsideBorder)

	w.ShowAndRun()

}
