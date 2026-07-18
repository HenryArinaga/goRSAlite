package screens

import (
	"fmt"
	"image/color"
	"log"

	"goRSAlite/internal/controller"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen(w fyne.Window, appController *controller.Controller) fyne.CanvasObject {
	var factorButtons *fyne.Container

	entry := widget.NewEntry()
	messageEntry := widget.NewEntry()
	entry.SetText(appController.CurrentInputText)
	messageEntry.Hide()
	entry.OnChanged = func(newText string) {
		appController.CurrentInputText = newText
	}

	titleLabel := container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("Prime Factorization"),
	)
	entry.SetPlaceHolder("Enter a postive integer...")

	gutter := canvas.NewRectangle(color.Transparent)
	gutter.SetMinSize(fyne.NewSize(200, 0))

	page := container.NewHBox(
		gutter,
		titleLabel,
	)
	progressBar := widget.NewProgressBarInfinite()
	progressBar.Hide()

	resultLabel := widget.NewLabel("The Factors are: ")
	resultLabel.Wrapping = fyne.TextWrapWord

	entry.Text = appController.CurrentInputText

	canceledLabel := widget.NewLabel("Operation Canceled!")
	canceledLabel.Hide()

	cancelButton := widget.NewButton("Cancel", func() {
		appController.CancelRunningFunction()
		progressBar.Hide()
		canceledLabel.Show()

	})
	cancelButton.Hide()

	RsaButton := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("RSA Demo", func() {
			go func() {
				appController.DoRsa(appController.Totient)
				fyne.Do(func() {
					if appController.BenchmarkOn == true {
						resultLabel.SetText(fmt.Sprintf(
							"RSA Values\nFactored number: %v \nP: %v, Q: %v\nEuler's Totient "+
								"%v\nPublic Exponent e: %v\nPrivate Exponent d: %v"+
								"\nPublic Key: (%v , %v)\nPrivate Key: (%v, %v)\nBenchmark Time: %v",
							appController.FactoredNumber,
							appController.LatestFactors[0],
							appController.LatestFactors[1],
							appController.Totient,
							appController.E,
							appController.D,
							appController.E,
							appController.LatestFactors[0]*appController.LatestFactors[1],
							appController.D,
							appController.LatestFactors[0]*appController.LatestFactors[1],
							appController.RsaDuration,
						))
					} else {
						resultLabel.SetText(fmt.Sprintf(
							"RSA Values\nFactored number: %v \nP: %v, Q: %v\nEuler's Totient "+
								"%v\nPublic Exponent e: %v\nPrivate Exponent d: %v"+
								"\nPublic Key: (%v , %v)\nPrivate Key: (%v, %v)",
							appController.FactoredNumber,
							appController.LatestFactors[0],
							appController.LatestFactors[1],
							appController.Totient,
							appController.E,
							appController.D,
							appController.E,
							appController.LatestFactors[0]*appController.LatestFactors[1],
							appController.D,
							appController.LatestFactors[0]*appController.LatestFactors[1],
						))
					}
				})
			}()
			messageEntry.Show()
		}))
	if len(appController.LatestFactors) == 2 {
		RsaButton.Show()
		RsaButton.Refresh()
	} else {
		RsaButton.Hide()

	}

	if appController.Running == false {
		cancelButton.Hide()
		progressBar.Hide()
	} else {
		cancelButton.Show()
		progressBar.Show()
	}

	if appController.RsaRunning == false {
		cancelButton.Hide()
		progressBar.Hide()
	} else {
		cancelButton.Show()
		progressBar.Show()
	}

	if len(appController.LatestFactors) == 0 {
		resultLabel.SetText("Factors will be displayed here:")
	} else {
		if appController.BenchmarkOn == true {
			resultLabel.SetText(fmt.Sprintf("Factored number: %v \nFactors: %v \nMethod: %s \nBenchmark Time: %v", appController.FactoredNumber, appController.LatestFactors, appController.SelectedMethod, appController.Duration))
			fmt.Printf("\nTime to find Factors: %s\n", appController.Duration)
		} else {
			resultLabel.SetText(fmt.Sprintf("Factored number: %v \nFactors: %v \nMethod: %s", appController.FactoredNumber, appController.LatestFactors, appController.SelectedMethod))
		}

	}
	// create the same buttons in the same container
	factorButtons = container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Factor", func() {
			text := entry.Text
			numberToFactor, err := strconv.Atoi(text)
			if appController.SelectedMethod == "" {
				dialog.ShowInformation("No Method Selected", "Pleaes selecte a Method", w)
				return
			}
			if err != nil {
				dialog.ShowInformation("Non-number characters entered", "Pleaes enter valid numbers", w)
				return
			}
			progressBar.Show()
			cancelButton.Show()
			appController.DoFactorization(numberToFactor)
			go func() {
				factors := <-appController.Done
				fyne.Do(func() {
					if appController.Running == false {
						cancelButton.Hide()
						progressBar.Hide()
					}
					if appController.BenchmarkOn == true {
						canceledLabel.Hide()
						resultLabel.SetText(fmt.Sprintf("Factored number: %v \nFactors: %v \nMethod: %s \nBenchmark Time: %v", appController.FactoredNumber, appController.LatestFactors, appController.SelectedMethod, appController.Duration))
						fmt.Printf("\nTime to find Factors: %s\n", appController.Duration)
					} else {
						canceledLabel.Hide()
						resultLabel.SetText(fmt.Sprintf("Factored number: %v \nFactors: %v \nMethod: %s", appController.FactoredNumber, factors, appController.SelectedMethod))
					}
					if len(factors) == 2 {
						RsaButton.Show()
						messageEntry.Show()
					} else {
						RsaButton.Hide()
						messageEntry.Hide()

					}
					factorButtons.Refresh()

				})
				log.Println(appController.FactoredNumber)
				log.Println(factors)
				log.Println("Current method:", appController.SelectedMethod)
			}()
		}),
		widget.NewButton("Clear", func() {
			entry.SetText("")
			appController.CurrentInputText = ("")
			RsaButton.Hide()
			messageEntry.Hide()
		}),
		cancelButton,
		RsaButton,
	)

	entry.Text = appController.CurrentInputText

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

	rectangle2Content := container.NewPadded(container.NewVBox(
		widget.NewLabel("Results"),
		container.NewVBox(resultLabel),
		container.NewVBox(canceledLabel),
		messageEntry,
	))

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
