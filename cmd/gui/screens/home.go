package screens

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen() fyne.CanvasObject {

	titleLabel := container.New(layout.NewVBoxLayout(), widget.NewLabel("Prime Factorization"))

	gutter := canvas.NewRectangle(color.Transparent)
	gutter.SetMinSize(fyne.NewSize(200, 0))

	page := container.NewHBox(
		gutter,
		titleLabel,
	)

	rectangle1 := canvas.NewRectangle(color.Black)
	rectangle1.SetMinSize(fyne.NewSize(500, 100))
	rectangle1Content := container.NewPadded(
		container.NewVBox(
			widget.NewLabel("Enter Number"),
			widget.NewEntry(),
			widget.NewButton("Factor", func() {}),
		),
	)

	panel := container.NewStack(
		rectangle1,
		rectangle1Content,
	)
	centeringContainer := container.New(layout.NewCenterLayout(), panel)
	insideWrapper := container.New(layout.NewVBoxLayout(), container.NewPadded(), page, centeringContainer)

	return insideWrapper
}
