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
	w := a.NewWindow("Hello World")
	text1 := canvas.NewText("menu 1", color.White)
	text2 := canvas.NewText("menu 2", color.White)

	w.SetContent(widget.NewLabel("Hello World!"))
	w.Resize(fyne.NewSize(1200, 800))

	HomeButton := widget.NewButton("Home", func() {
		log.Println("tapped")
	})

	leftBox := container.New(layout.NewVBoxLayout(), text1, text2, HomeButton)

	w.SetContent(container.New(layout.NewHBoxLayout(), leftBox, HomeButton))

	w.ShowAndRun()

}
