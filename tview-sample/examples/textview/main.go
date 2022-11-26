package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const corporate = `アジャイル フレームワークを活用して、高レベルの概要の堅牢な概要を提供します。 企業戦略への反復的なアプローチは、全体的な価値提案を促進するための共同思考を促進します。 職場の多様性とエンパワーメントを通じて、破壊的イノベーションの全体論的な世界観を有機的に成長させます。

有利なサバイバル戦略をテーブルにもたらし、積極的な支配を確実にします。 結局のところ、ジェネレーション X から進化したニューノーマルは、合理化されたクラウド ソリューションに向かって滑走路上にあります。 ユーザーがリアルタイムで作成したコンテンツには、オフショアリングのための複数のタッチポイントがあります。

簡単に達成できる成果を利用して、ベータ テストに大まかな付加価値のあるアクティビティを特定します。 DevOps からの追加のクリックスルーでデジタル デバイドをオーバーライドします。 情報ハイウェイに沿ったナノテクノロジーへの没入は、利益のみに焦点を当てるというループを閉じます。

[yellow] Enter キーを押してから、Tab/BackTab を押して単語を選択します`

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	numSelections := 0
	go func() {
		for _, word := range strings.Split(corporate, "") {
			if word == "は" {
				word = "[red]は[white]"
			}
			if word == "を" {
				word = fmt.Sprintf(`["%d"]を[""]`, numSelections)
				numSelections++
			}
			fmt.Fprintf(textView, "%s", word)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	textView.SetDoneFunc(func(key tcell.Key) {
		currentSelection := textView.GetHighlights()
		if key == tcell.KeyEnter {
			if len(currentSelection) > 0 {
				textView.Highlight()
			} else {
				textView.Highlight("0").ScrollToHighlight()
			}
		} else if len(currentSelection) > 0 {
			index, _ := strconv.Atoi(currentSelection[0])
			if key == tcell.KeyTab {
				index = (index + 1) % numSelections
			} else if key == tcell.KeyBacktab {
				index = (index - 1 + numSelections) % numSelections
			} else {
				return
			}
			textView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
		}
	})
	textView.SetBorder(true)
	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}
}
