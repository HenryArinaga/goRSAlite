package screens

import (
	"goRSAlite/internal/controller"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SettingsScreen(appController *controller.Controller) fyne.CanvasObject {

	benchmarkModeCheck := widget.NewCheck("Benchmark mode",
		func(turnOnBenchmarkMode bool) {
			appController.BenchmarkOn = turnOnBenchmarkMode

		})

	if appController.BenchmarkOn == true {
		benchmarkModeCheck.SetChecked(true)
	}
	bottomTextBenchmark := widget.NewLabel("Time your factor results")
	bottomTextBenchmark.TextStyle.Italic = true

	sieveModeCheck := widget.NewCheck("Sieve mode",
		func(turnOnSieveMode bool) {
			appController.SieveOn = turnOnSieveMode
		})

	if appController.SieveOn == true {
		sieveModeCheck.SetChecked(true)
	}

	bottomTextSieve := widget.NewLabel("Apply Sieve to find factors faster; Only usable with trial division")
	bottomTextSieve.TextStyle.Italic = true

	simdModeCheck := widget.NewCheck("SIMD Acceleration mode (coming soon)",
		func(turnOnSimdMode bool) {
			appController.SimdOn = turnOnSimdMode
		})
	recAppearance1 := canvas.NewRectangle(color.Black)
	recAppearance1.SetMinSize(fyne.NewSize(500, 50))

	recAppearance2 := canvas.NewRectangle(color.Black)
	recAppearance2.SetMinSize(fyne.NewSize(500, 50))

	recAppearance3 := canvas.NewRectangle(color.Black)
	recAppearance3.SetMinSize(fyne.NewSize(500, 50))

	recAppearance4 := canvas.NewRectangle(color.Black)
	recAppearance4.SetMinSize(fyne.NewSize(500, 50))
	simdModeCheck.Disable()

	optimizationSettings := widget.NewLabel("Optimization Settings")
	optimizationSettings.TextStyle = fyne.TextStyle{Bold: true}
	optimizationHeading := container.New(layout.NewCenterLayout(), optimizationSettings, layout.NewSpacer())

	benchmarkSubText := container.New(layout.NewVBoxLayout(), benchmarkModeCheck, bottomTextBenchmark)
	benchmarkModeCheckRow := container.New(layout.NewStackLayout(), recAppearance2, benchmarkSubText)
	benchmarkBox := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), optimizationHeading, benchmarkModeCheckRow, layout.NewSpacer())

	sieveSubText := container.New(layout.NewVBoxLayout(), sieveModeCheck, bottomTextSieve)
	sieveRow := container.New(layout.NewStackLayout(), recAppearance3, sieveSubText, layout.NewSpacer())
	sieveBox := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), sieveRow, layout.NewSpacer())

	simdRow := container.New(layout.NewStackLayout(), recAppearance4, simdModeCheck)
	simdBox := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), simdRow, layout.NewSpacer())

	centered1 := container.New(layout.NewVBoxLayout() /*appearanceBox,*/, benchmarkBox, sieveBox, simdBox, layout.NewSpacer())

	limitSize := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), centered1, layout.NewSpacer())
	return limitSize
}
