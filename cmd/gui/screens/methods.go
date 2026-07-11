package screens

import (
	"image/color"
	"log"

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
	cardList []*MethodCard
}

var selectedMethod string

func (methodCardClicked *MethodCard) Tapped(eventInfo *fyne.PointEvent) {

	for _, card := range methodCardClicked.cardList {
		card.Selected = false
		card.Rectangle.FillColor = color.Black
		card.Refresh()

	}
	selectedMethod = methodCardClicked.MethodName
	methodCardClicked.Selected = true
	methodCardClicked.Rectangle.FillColor = color.RGBA{R: 80, G: 140, B: 255, A: 255}
	methodCardClicked.Refresh()

	selectedMethod = methodCardClicked.MethodName
	log.Println("Select set to", methodCardClicked.MethodName)

}

func (methodCardClicked *MethodCard) CreateRenderer() fyne.WidgetRenderer {

	widgetContainer := container.New(
		layout.NewStackLayout(),
		methodCardClicked.Rectangle,
		methodCardClicked.Label,
	)

	return widget.NewSimpleRenderer(widgetContainer)
}

func MethodScreen() fyne.CanvasObject {

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
		MethodName: "Fermat Factorization",
		Rectangle:  rectangle3,
		Label:      fermat,
	}
	card3.ExtendBaseWidget(&card3)

	var cardList []*MethodCard = []*MethodCard{&card1, &card2, &card3}
	card1.cardList = cardList
	card2.cardList = cardList
	card3.cardList = cardList

	for _, v := range cardList {
		if selectedMethod == v.MethodName {
			v.Rectangle.FillColor = color.RGBA{R: 80, G: 140, B: 255, A: 255}
			v.Selected = true
		} else {
			v.Rectangle.FillColor = color.Black
			v.Selected = false
		}

	}

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
