package screens

import (
	"fmt"
	"image/color"
	"log"
	"time"

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
	messageEntry := widget.NewEntry()
	resultEncDecLabel := widget.NewLabel("Encrypted text will appear here")
	resultEncDecLabel.Wrapping = fyne.TextWrapWord
	resultEncDecScroll := container.NewVScroll(resultEncDecLabel)
	resultEncDecScroll.SetMinSize(fyne.NewSize(760, 150))
	resultEncDecLabel.Hide()

	resultLabel := widget.NewLabel("The Factors are: ")
	resultLabel.Wrapping = fyne.TextWrapWord

	encDecButton := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Encrypt / Decrypt", func() {
			input := messageEntry.Text
			go func() {
				appController.DoEncrypt(input)
				EncryptedText := <-appController.EncryptionDone
				appController.DoDecrypt(appController.CurrentInputText)
				DecryptedText := <-appController.DecryptionDone
				fyne.Do(func() {
					for i := 0; i < len(messageEntry.Text); i++ {
						if int(messageEntry.Text[i]) >= appController.FactoredNumber || int(messageEntry.Text[i]) <= 0 {
							fmt.Printf("Message byte %d must satisfy 0 < byte < n for RSA encryption.\n", messageEntry.Text[i])
							messageLenWarning := fmt.Sprintf("The number you have chosen for the RSA demo is too small for encryption/decryption for the message you want to use ")
							resultLabel.SetText(messageLenWarning)
							return
						}
					}
					resultText := fmt.Sprintf("Encrypted Text: %v\nDecrypted Text: %c\nNumber of decrypted bytes: %d", EncryptedText, DecryptedText, len(DecryptedText))
					resultEncDecLabel.Show()
					resultEncDecLabel.SetText(resultText)
					appController.History = append(appController.History, controller.LogEntry{
						Time:        time.Now(),
						Kind:        "Encrypt / Decrypt",
						Input:       fmt.Sprintf("%d", appController.FactoredNumber),
						Result:      resultText,
						Method:      appController.SelectedMethod,
						Duration:    appController.EncryptionDuration,
						Sieve:       appController.SieveOn,
						SIMD:        appController.SimdOn,
						ExtendedGcd: appController.ExtendedGcdOn,
					})

				})

				log.Println(EncryptedText)
			}()

		}))

	entry := widget.NewEntry()
	entry.SetText(appController.CurrentInputText)
	messageEntry.SetPlaceHolder("Enter a short message to encrpyt")
	messageEntry.SetText(appController.CurrentMessageText)
	messageEntry.Hide()
	encDecButton.Hide()
	entry.OnChanged = func(newText string) {
		appController.CurrentInputText = newText
	}

	messageEntry.OnChanged = func(newMessageText string) {
		appController.CurrentMessageText = newMessageText
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
			progressBar.Show()
			messageEntry.Show()
			encDecButton.Show()
			go func() {
				appController.DoRsa()
				result := <-appController.RsaDone
				fyne.Do(func() {
					progressBar.Hide()
					resultText := fmt.Sprintf(
						"RSA Values\nFactored number: %v \nP: %v, Q: %v\nEuler's Totient "+
							"%v\nPublic Exponent e: %v\nPrivate Exponent d: %v"+
							"\nPublic Key: (%v , %v)\nPrivate Key: (%v, %v)",
						appController.FactoredNumber,
						appController.LatestFactors[0],
						appController.LatestFactors[1],
						result.Totient,
						result.PublicExponent,
						result.PrivateExponent,
						result.PublicExponent,
						appController.LatestFactors[0]*appController.LatestFactors[1],
						result.PrivateExponent,
						appController.LatestFactors[0]*appController.LatestFactors[1],
					)
					if appController.BenchmarkOn {
						resultText += fmt.Sprintf("\nBenchmark Time: %v", appController.RsaDuration)
					}
					resultLabel.SetText(resultText)
					appController.History = append(appController.History, controller.LogEntry{
						Time:        time.Now(),
						Kind:        "RSA Demo",
						Input:       fmt.Sprintf("%d", appController.FactoredNumber),
						Result:      resultText,
						Method:      appController.SelectedMethod,
						Duration:    appController.RsaDuration,
						Sieve:       appController.SieveOn,
						SIMD:        appController.SimdOn,
						ExtendedGcd: appController.ExtendedGcdOn,
					})
				})
			}()

		}))

	// create the same buttons in the same container
	factorButtons = container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Factor", func() {
			messageEntry.Hide()
			encDecButton.Hide()
			appController.CurrentMessageText = ("")
			messageEntry.SetText("")
			resultEncDecLabel.Hide()
			text := entry.Text
			numberToFactor, err := strconv.Atoi(text)
			if appController.SelectedMethod == "" {
				dialog.ShowInformation("No Method Selected", "Please select a Method", w)
				return
			}
			if err != nil {
				dialog.ShowInformation("Non-number characters entered", "Please enter valid numbers", w)
				return
			}
			progressBar.Show()
			cancelButton.Show()
			go func() {
				appController.DoFactorization(numberToFactor)
				factors := <-appController.Done
				fyne.Do(func() {
					if appController.Running == false {
						cancelButton.Hide()
						progressBar.Hide()
					}
					if appController.BenchmarkOn == true {
						canceledLabel.Hide()
						resultText := fmt.Sprintf("Factors: %v", factors)
						displayText := fmt.Sprintf(
							"Factored number: %v \n%s \nMethod: %s \nBenchmark Time: %v",
							appController.FactoredNumber,
							resultText,
							appController.SelectedMethod,
							appController.Duration,
						)
						if len(factors) != 2 {
							displayText += fmt.Sprintf("\nRSA Demo requires exactly two prime factors")

						}
						resultLabel.SetText(displayText)
						appController.History = append(appController.History, controller.LogEntry{
							Time:        time.Now(),
							Kind:        "Factorization",
							Input:       text,
							Result:      resultText,
							Method:      appController.SelectedMethod,
							Duration:    appController.Duration,
							Sieve:       appController.SieveOn,
							SIMD:        appController.SimdOn,
							ExtendedGcd: appController.ExtendedGcdOn,
						})
						fmt.Printf("\nTime to find Factors: %s\n", appController.Duration)
					} else {
						canceledLabel.Hide()
						resultText := fmt.Sprintf("Factors: %v", factors)
						displayText := fmt.Sprintf(
							"Factored number: %v \n%s \nMethod: %s",
							appController.FactoredNumber,
							resultText,
							appController.SelectedMethod,
						)
						if len(factors) != 2 {
							displayText += fmt.Sprintf("\nRSA Demo requires exactly two prime factors")

						}
						resultLabel.SetText(displayText)
						appController.History = append(appController.History, controller.LogEntry{
							Time:        time.Now(),
							Kind:        "Factorization",
							Input:       text,
							Result:      resultText,
							Method:      appController.SelectedMethod,
							Duration:    appController.Duration,
							Sieve:       appController.SieveOn,
							SIMD:        appController.SimdOn,
							ExtendedGcd: appController.ExtendedGcdOn,
						})
					}
					if len(factors) == 2 {
						RsaButton.Show()
					} else {
						RsaButton.Hide()
						messageEntry.Hide()
						encDecButton.Hide()
						resultEncDecLabel.Hide()
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
			messageEntry.SetText("")
			RsaButton.Hide()
			messageEntry.Hide()
			encDecButton.Hide()
			appController.CurrentMessageText = ("")
			messageEntry.SetText("")
			resultEncDecLabel.Hide()
			resultLabel.SetText("Factors will be displayed here:")

		}),
		cancelButton,
		RsaButton,
	)

	if len(appController.LatestFactors) == 2 {
		RsaButton.Show()
		RsaButton.Refresh()
	} else {
		appController.CurrentMessageText = ("")
		messageEntry.SetText("")
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
		encDecButton,
		resultEncDecScroll,
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
