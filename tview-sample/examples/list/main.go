package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("リストアイテム (1)", "雪谷高校", 'a', nil).
		AddItem("リストアイテム (2)", "大森学園", 'b', nil).
		AddItem("リストアイテム (3)", "工科高校", 'c', nil).
		AddItem("リストアイテム (4)", "晴海総合高校", 'd', nil).
		AddItem("終了", "押すと終了します", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
