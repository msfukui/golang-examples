package main

import (
	"time"

	"github.com/rivo/tview"
)

const refreshInterval = 500 * time.Millisecond

func currentTimeString() string {
	t := time.Now()
	return t.Format("現在の時刻は 15:04:05 です.")
}

func updateTime(app *tview.Application, view *tview.Modal) string {
	for {
		time.Sleep(refreshInterval)
		app.QueueUpdateDraw(func() {
			view.SetText(currentTimeString())
		})
	}
}

func main() {
	app := tview.NewApplication()
	view := tview.NewModal().
		SetText(currentTimeString()).
		AddButtons([]string{"終了"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "終了" {
				app.Stop()
			}
		})

	go updateTime(app, view)
	if err := app.SetRoot(view, false).Run(); err != nil {
		panic(err)
	}
}
