package screens

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SettingsScreen() fyne.CanvasObject {

	lightModeCheck := widget.NewCheck("Light mode (coming soon)",
		func(turnOnLightMode bool) {

		})

	benchmarkModeCheck := widget.NewCheck("Benchmark mode",
		func(turnOnBenchmarkMode bool) {
			BenchmarkOn = turnOnBenchmarkMode

		})

	if BenchmarkOn == true {
		benchmarkModeCheck.SetChecked(true)
	}

	sieveModeCheck := widget.NewCheck("Sieve mode (coming soon)",
		func(turnOnSieveMode bool) {

		})

	simdModeCheck := widget.NewCheck("SIMD Acceleration mode (coming soon)",
		func(turnOnSimdMode bool) {

		})
	recAppearance1 := canvas.NewRectangle(color.Black)
	recAppearance1.SetMinSize(fyne.NewSize(500, 50))
	lightModeCheck.Disable()

	recAppearance2 := canvas.NewRectangle(color.Black)
	recAppearance2.SetMinSize(fyne.NewSize(500, 50))

	recAppearance3 := canvas.NewRectangle(color.Black)
	recAppearance3.SetMinSize(fyne.NewSize(500, 50))
	sieveModeCheck.Disable()

	recAppearance4 := canvas.NewRectangle(color.Black)
	recAppearance4.SetMinSize(fyne.NewSize(500, 50))
	simdModeCheck.Disable()

	Appearance := widget.NewLabel("Appearance")
	Appearance.TextStyle = fyne.TextStyle{Bold: true}
	appearanceHeading := container.New(layout.NewCenterLayout(), Appearance, layout.NewSpacer())
	lightModeRow := container.New(layout.NewStackLayout(), recAppearance1, layout.NewSpacer(), lightModeCheck)
	appearanceBox := container.New(layout.NewVBoxLayout(), appearanceHeading, lightModeRow, layout.NewSpacer())

	optimizationSettings := widget.NewLabel("Optimization Settings")
	optimizationSettings.TextStyle = fyne.TextStyle{Bold: true}
	optimizationHeading := container.New(layout.NewCenterLayout(), optimizationSettings, layout.NewSpacer())
	optimizationRow := container.New(layout.NewStackLayout(), recAppearance2, benchmarkModeCheck)
	optimizationBox := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), optimizationHeading, optimizationRow, layout.NewSpacer())

	sieveRow := container.New(layout.NewStackLayout(), recAppearance3, sieveModeCheck)
	sieveBox := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), sieveRow, layout.NewSpacer())

	simdRow := container.New(layout.NewStackLayout(), recAppearance4, simdModeCheck)
	simdBox := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), simdRow, layout.NewSpacer())

	centered1 := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), appearanceBox, optimizationBox, sieveBox, simdBox, layout.NewSpacer())

	limitSize := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), centered1, layout.NewSpacer())
	return limitSize
}
