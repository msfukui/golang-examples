package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const refreshInterval = 500 * time.Millisecond

func drawTime(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	timeStr := time.Now().Format("現在の時刻は 15:04:05 です.")
	tview.Print(screen, timeStr, x, height/2, width, tview.AlignCenter, tcell.ColorLime)
	return 0, 0, 0, 0
}

func refresh(app *tview.Application) string {
	tick := time.NewTicker(refreshInterval)
	for {
		<-tick.C
		app.Draw()
	}
}

func main() {
	app := tview.NewApplication()
	view := tview.NewBox().SetDrawFunc(drawTime)

	go refresh(app)
	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}
