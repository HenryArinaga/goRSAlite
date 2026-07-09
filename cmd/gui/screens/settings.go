package screens

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MethodCard struct {
	MethodName string
	Rectangle  *canvas.Rectangle
	Label      *widget.Label
	Selected   bool
	widget.BaseWidget
}

var selectedMethod string

func (methodCardClicked *MethodCard) Tapped(eventInfo *fyne.PointEvent) {

	selectedMethod = methodCardClicked.MethodName
	methodCardClicked.Selected = true
	methodCardClicked.Rectangle.FillColor = color.RGBA{R: 80, G: 140, B: 255, A: 255}
	methodCardClicked.Refresh()

}

func (methodCardClicked *MethodCard) TappedSecondary(eventInfo *fyne.PointEvent) {

}

func (methodCardClicked *MethodCard) CreateRenderer() fyne.WidgetRenderer {

	widgetContainer := container.New(
		layout.NewStackLayout(),
		methodCardClicked.Rectangle,
		methodCardClicked.Label,
	)

	return widget.NewSimpleRenderer(widgetContainer)
}

func SettingsScreen() fyne.CanvasObject {

	rectangle1 := canvas.NewRectangle(color.Black)
	rectangle1.SetMinSize(fyne.NewSize(500, 50))
	rectangle2 := canvas.NewRectangle(color.Black)
	rectangle2.SetMinSize(fyne.NewSize(500, 50))
	rectangle3 := canvas.NewRectangle(color.Black)
	rectangle3.SetMinSize(fyne.NewSize(500, 50))
	trialDivison := widget.NewLabel("Trial Division")
	sqrt := widget.NewLabel("Square Root")
	fermat := widget.NewLabel("Fermat Method")

	card1 := MethodCard{
		MethodName: "Trial Division",
		Rectangle:  rectangle1,
		Label:      trialDivison,
	}
	card1.ExtendBaseWidget(&card1)

	card2 := MethodCard{
		MethodName: "Square Root",
		Rectangle:  rectangle2,
		Label:      sqrt,
	}
	card2.ExtendBaseWidget(&card2)

	card3 := MethodCard{
		MethodName: "Fermat Method",
		Rectangle:  rectangle3,
		Label:      fermat,
	}
	card3.ExtendBaseWidget(&card3)

	content1 := container.New(layout.NewStackLayout(), &card1)
	content2 := container.New(layout.NewStackLayout(), &card2)
	content3 := container.New(layout.NewStackLayout(), &card3)

	centered1 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), content1, layout.NewSpacer())
	centered2 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), content2, layout.NewSpacer())
	centered3 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), content3, layout.NewSpacer())

	cards := container.New(layout.NewVBoxLayout(), centered1, centered2, centered3)

	centeredCards := container.New(layout.NewCenterLayout(), cards)

	return centeredCards
}
