package main

import (
	"goRSAlite/cmd/gui/screens"
	"goRSAlite/internal/controller"
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

	appController := &controller.Controller{Done: make(chan []int), EncryptionDone: make(chan []int), DecryptionDone: make(chan []int)}

	homescreen := screens.HomeScreen(w, appController)

	centerTop := container.New(layout.NewCenterLayout(), canvas.NewText("Welcome to the RSA Factorization Tool!", color.White))

	rightTop := container.New(layout.NewHBoxLayout(), widget.NewLabel("Right Top"))

	topBar := container.NewBorder(nil, nil, nil, rightTop, centerTop)

	centerContent := container.New(layout.NewStackLayout(), homescreen)

	HomeButton := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		log.Println("tapped")
		centerContent.Objects = []fyne.CanvasObject{homescreen}
		centerContent.Refresh()
	})
	MethodsButton := widget.NewButtonWithIcon("Methods", theme.ComputerIcon(), func() {
		log.Println("Methods tapped")
		centerContent.Objects = []fyne.CanvasObject{screens.MethodScreen(appController)}
		centerContent.Refresh()
	})
	SettingsButton := widget.NewButton("Settings", func() {
		log.Println("Settings tapped")
		centerContent.Objects = []fyne.CanvasObject{screens.SettingsScreen(appController)}
		centerContent.Refresh()
	})
	LogsButton := widget.NewButton("Logs", func() {
		log.Println("Logs tapped")
		centerContent.Objects = []fyne.CanvasObject{screens.LogsScreen(appController)}
		centerContent.Refresh()
	})

	w.Resize(fyne.NewSize(1200, 800))

	leftSidebar := container.NewVBox(
		HomeButton,
		MethodsButton,
		LogsButton,
		SettingsButton,
	)

	outsideBorder := container.NewBorder(topBar, nil, leftSidebar, nil, centerContent)

	w.SetContent(outsideBorder)

	w.ShowAndRun()

}
