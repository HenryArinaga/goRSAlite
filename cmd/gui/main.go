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

	appController := &controller.Controller{Done: make(chan []int), EncryptionDone: make(chan []int), DecryptionDone: make(chan []int), RsaDone: make(chan controller.RSAResult)}

	homeScreen := screens.HomeScreen(w, appController)

	logsScreen, refreshLogs := screens.LogsScreen(w, appController)

	centerTop := container.New(layout.NewCenterLayout(), canvas.NewText("Welcome to the RSA Factorization Tool!", color.White))

	topBar := container.NewBorder(nil, nil, nil, nil, centerTop)

	centerContent := container.New(layout.NewStackLayout(), homeScreen)

	HomeButton := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		log.Println("tapped")
		centerContent.Objects = []fyne.CanvasObject{homeScreen}
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
		refreshLogs()
		log.Println("Logs tapped")
		centerContent.Objects = []fyne.CanvasObject{logsScreen}
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
