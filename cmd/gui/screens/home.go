package screens

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"goRSAlite/internal/factorization"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var BenchmarkOn bool

func HomeScreen(w fyne.Window) fyne.CanvasObject {

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
	progressBar := widget.NewProgressBarInfinite()
	progressBar.Hide()

	resultLabel := widget.NewLabel("The Factors are: ")

	// create the same buttons in the same container
	factorButtons := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Factor", func() {
			text := entry.Text
			numberToFactor, err := strconv.Atoi(text)
			if selectedMethod == "" {
				dialog.ShowInformation("No Method Selected", "Pleaes selecte a Method", w)
				return
			}
			if err != nil {
				dialog.ShowInformation("Non-number characters entered", "Pleaes enter valid numbers", w)
				return
			}
			factorResult := make(chan []int, 1)
			switch selectedMethod {
			case "Trial Division":
				start := time.Now()
				progressBar.Show()
				go func() {
					factorization.FactorTrialDivision(numberToFactor, factorResult)
					factors := <-factorResult
					duration := time.Since(start)
					fyne.Do(func() {
						progressBar.Hide()
						if BenchmarkOn == true {
							resultLabel.SetText(fmt.Sprintf("Factors: %v \nMethod: Trial Division\nBenchmark Time: %v", factors, duration))
							fmt.Printf("\nTime to find Factors: %s\n", duration)
						} else {
							resultLabel.SetText(fmt.Sprintf("Factors: %v \nMethod: Trial Division", factors))
						}
					})
					log.Println(factors)
					log.Println("Current method:", selectedMethod)
				}()

			case "Square Root":
				start := time.Now()
				progressBar.Show()
				go func() {
					factorization.FactorSqrt(numberToFactor, factorResult)
					factors := <-factorResult
					duration := time.Since(start)
					fyne.Do(func() {
						progressBar.Hide()
						if BenchmarkOn == true {
							resultLabel.SetText(fmt.Sprintf("Factors: %v\nMethod: Square Root\nBenchmark Time: %v", factors, duration))
						} else {
							resultLabel.SetText(fmt.Sprintf("Factors: %v\nMethod: Square Root", factors))
						}
					})
					log.Println(factors)
					log.Println("Current method:", selectedMethod)
				}()

			case "Fermat Factorization":
				start := time.Now()
				progressBar.Show()
				go func() {
					factors := factorization.FactorFermat(numberToFactor)
					factorResult <- factors

					duration := time.Since(start)
					fyne.Do(func() {
						progressBar.Hide()
						if BenchmarkOn == true {
							resultLabel.SetText(fmt.Sprintf("Factors: %v\nMethod: Fermat\nBenchmark Time: %v", factors, duration))
						} else {
							resultLabel.SetText(fmt.Sprintf("Factors: %v\nMethod: Fermat", factors))
						}
					})
					log.Println(factors)
					log.Println("Current method:", selectedMethod)
				}()
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
		progressBar,
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
