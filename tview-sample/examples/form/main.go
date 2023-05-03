package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddDropDown("タイトル", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("氏名（姓）", "", 20, nil, nil).
		AddInputField("氏名（名）", "", 20, nil, nil).
		AddTextArea("住所", "", 40, 0, 0, nil).
		AddCheckbox("年齢(18+)", false, nil).
		AddPasswordField("パスワード", "", 10, '*', nil).
		AddButton("保存", nil).
		AddButton("終了", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("いくつか情報を入力してください").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
