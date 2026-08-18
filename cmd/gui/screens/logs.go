package screens

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"goRSAlite/internal/controller"
	"log"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Result struct {
	Time        string `json:"Time"`
	Kind        string `json:"Kind"`
	Input       string `json:"Input"`
	Result      string `json:"Result"`
	Method      string `json:"Method"`
	Duration    string `json:"Duration"`
	Sieve       bool   `json:"Sieve"`
	SIMD        bool   `json:"SIMD"`
	ExtendedGcd bool   `json:"ExtendedGcd"`
}

func LogsScreen(w fyne.Window, appController *controller.Controller) (fyne.CanvasObject, func()) {

	var saveAlert *widget.PopUp
	saveAlertBox := container.NewVBox(
		widget.NewLabel("Operations Saved Successfully"),
		widget.NewButton("Close", func() {
			saveAlert.Hide()
		}),
	)

	logsHeading1 := widget.NewLabel("Operation History")
	logsHeading1.TextStyle = fyne.TextStyle{Bold: true}
	logsCentered := container.New(layout.NewCenterLayout(), logsHeading1)
	display := container.New(layout.NewVBoxLayout(), logsCentered)

	content := container.NewVBox()
	scroll := container.NewVScroll(content)
	refreshLogs := func() {
		content.RemoveAll()
		for _, entry := range appController.History {
			text := fmt.Sprintf(
				"Time: %s\nType: %s\nInput: %s\nResult: %s\nMethod: %s\nDuration: %s\nSieve: %t\nSIMD: %t\nExtended GCD: %t\n",
				entry.Time.Format("15:04:05"),
				entry.Kind,
				entry.Input,
				entry.Result,
				entry.Method,
				entry.Duration,
				entry.Sieve,
				entry.SIMD,
				entry.ExtendedGcd,
			)
			label := widget.NewLabel(text)
			label.Wrapping = fyne.TextWrapWord
			content.Add(label)
		}
		content.Refresh()
		scroll.ScrollToOffset(appController.LogsScrollOffset)
	}

	scroll.SetMinSize(fyne.NewSize(500, 500))

	scroll.OnScrolled = func(pos fyne.Position) {
		appController.LogsScrollOffset = pos
	}
	clearButton := container.New(layout.NewHBoxLayout(),
		widget.NewButton("Clear", func() {
			appController.History = nil
			refreshLogs()
		}),
	)

	CsvButton := container.New(layout.NewHBoxLayout(),
		widget.NewButton("Export CSV", func() {
			file, err := os.Create("output.csv")
			if err != nil {
				dialog.ShowInformation("CSV ERROR", "CSV Save Failed", w)
				log.Fatal(err)
			}
			defer file.Close()

			// Create a CSV writer
			writer := csv.NewWriter(file)
			defer writer.Flush()

			// Write header
			header := []string{"Time", "Kind", "Input", "Result", "Method", "Duration", "Sieve", "SIMD", "ExtendedGcd"}
			if err := writer.Write(header); err != nil {
				log.Fatal(err)
			}

			// Write data rows
			records := [][]string{}
			for _, entry := range appController.History {
				row := []string{entry.Time.Format("2006-01-02 15:04:05"),
					entry.Kind,
					entry.Input,
					entry.Result,
					entry.Method,
					entry.Duration.String(),
					strconv.FormatBool(entry.Sieve),
					strconv.FormatBool(entry.SIMD),
					strconv.FormatBool(entry.ExtendedGcd)}
				records = append(records, row)
			}
			if err := writer.WriteAll(records); err != nil {
				log.Fatal(err)
			}
			saveAlert = widget.NewModalPopUp(saveAlertBox, w.Canvas())

			saveAlert.Show()
		}),
	)

	JsonButton := container.New(layout.NewHBoxLayout(),
		widget.NewButton("Export JSON", func() {
			results := []Result{}
			for _, entry := range appController.History {
				result := Result{
					Time:        entry.Time.Format("2006-01-02 15:04:05"),
					Kind:        entry.Kind,
					Input:       entry.Input,
					Result:      entry.Result,
					Method:      entry.Method,
					Duration:    entry.Duration.String(),
					Sieve:       entry.Sieve,
					SIMD:        entry.SIMD,
					ExtendedGcd: entry.ExtendedGcd,
				}
				results = append(results, result)
			}

			jsonData, err := json.MarshalIndent(results, "", "  ")
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				return
			}
			err = os.WriteFile("output.json", jsonData, 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				dialog.ShowInformation("JSON ERROR", "JSON Save Failed", w)
				return
			}

			fmt.Println("JSON file exported successfully.")
			saveAlert = widget.NewModalPopUp(saveAlertBox, w.Canvas())
			saveAlert.Show()
		}),
	)

	topRow := container.NewBorder(nil, nil, nil, clearButton, container.NewHBox(CsvButton, JsonButton))
	centered := container.New(layout.NewVBoxLayout(), display, scroll, topRow, layout.NewSpacer())

	return centered, refreshLogs
}
