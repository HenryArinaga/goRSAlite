package screens

import (
	"fmt"
	"image/color"
	"log"

	"goRSAlite/internal/factorization"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen() fyne.CanvasObject {

	entry := widget.NewEntry()
	titleLabel := container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("Prime Factorization"),
	)

	gutter := canvas.NewRectangle(color.Transparent)
	gutter.SetMinSize(fyne.NewSize(200, 0))

	page := container.NewHBox(
		gutter,
		titleLabel,
	)

	resultLabel := widget.NewLabel("The Factors are: ")

	// create the same buttons in the same container
	factorButtons := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Factor", func() {
			text := entry.Text
			numberToFactor, err := strconv.Atoi(text)
			if err != nil {
				return
			}

			switch selectedMethod {

			case "Trial Division":
				factorResult := factorization.FactorTrialDivision(numberToFactor)
				resultLabel.SetText(fmt.Sprintf("Factors: %v", factorResult))
				log.Println(factorResult)
				log.Println("Current method:", selectedMethod)
			case "Square Root":
				factorResult := factorization.FactorSqrt(numberToFactor)
				resultLabel.SetText(fmt.Sprintf("Factors: %v", factorResult))
				log.Println(factorResult)
				log.Println("Current method:", selectedMethod)
			case "Fermat Factorization":
				factorResult := factorization.FactorFermat(numberToFactor)
				resultLabel.SetText(fmt.Sprintf("Factors: %v", factorResult))
				log.Println(factorResult)
				log.Println("Current method:", selectedMethod)
			default:
				fmt.Println("Invalid method selected")
			}

		}),
		widget.NewButton("Clear", func() {
			entry.SetText("")
		}),
	)

	//pass buttons to rectangle container so it sits on top of the rectangle
	rectangle1 := canvas.NewRectangle(color.Black)
	rectangle1.SetMinSize(fyne.NewSize(500, 100))

	rectangle1Content := container.NewPadded(
		container.NewVBox(
			widget.NewLabel("Enter Number"),
			entry,
			factorButtons,
		),
	)

	panel := container.NewStack(
		rectangle1,
		rectangle1Content,
	)

	rectangle2 := canvas.NewRectangle(color.Black)
	rectangle2.SetMinSize(fyne.NewSize(800, 400))

	rectangle2Content := container.NewPadded(
		container.NewVBox(
			widget.NewLabel("Results"),
			container.NewVBox(
				resultLabel),
		),
	)

	outputPanel := container.NewStack(
		rectangle2,
		rectangle2Content,
	)

	gap := canvas.NewRectangle(color.Transparent)
	gap.SetMinSize(fyne.NewSize(0, 16))

	verticalContainer := container.New(
		layout.NewVBoxLayout(),
		page,
		panel,
		gap,
		outputPanel,
	)

	centeringContainer := container.New(
		layout.NewCenterLayout(),
		container.NewPadded(),
		verticalContainer,
	)

	insideWrapper := container.New(
		layout.NewVBoxLayout(),
		container.NewPadded(),
		centeringContainer,
	)

	return insideWrapper
}
