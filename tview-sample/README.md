# tview-sample

A sample codes in use `tview`.

https://github.com/rivo/tview

WSL2 + Windows Terminal でサンプルコードが盛大に文字化けしたため一旦見送り。
→ bubble tea と同じく RUNEWIDTH_EASTASIAN の設定でおおよそ回避可能なことを確認。
過去の記事などでは非ascii文字対応していない、という記載もあったが、2022年11月時点では一旦考慮した実装がされている模様。
ただし、Ambiguous な文字表示を2文字幅と判定する環境は考慮されていなさそうなため、 bubble tea の方がその点一貫している。

* 以下、調べた結果のメモ

    * 結論: 実行前に RUNEWIDTH_EASTASIAN=0 を設定することで一旦回避は可能

        example: $ RUNEWIDTH_EASTASIAN=0 go run hello.go

    * 原因: East Asian Width における Ambiguous の扱いのずれの問題

    * Windows Terminal ではすべて半角文字扱いで強制される仕様の模様

        https://github.com/microsoft/terminal/pull/2928
        https://qiita.com/good_kobe/items/2da02b8d141984d1bc4d

    * tview が依存している tcell が文字幅の判定に go-runewidth を使っている

        https://github.com/gdamore/tcell
        https://github.com/mattn/go-runewidth

    * go-runewidth は既定ではlocaleを見てcjkだと罫線については全角文字として判定している

    * 表示幅は全角として判定しているのに実際の表示は半角なので半分でしか表示されない事象が発生する

    * go-runewidth では環境変数 RUNEWIDTH_EASTASIAN に 1 以外を設定することで、強制的に半角文字として判定させることが可能

        c.f. https://noborus.github.io/blog/runewidth/

    * 一方で tview が依存している uniseg でも同様に文字幅を判定する関数が存在する

        https://github.com/rivo/uniseg

    * こちらは Ambiguous, Natural はいずれも半角文字(文字幅1)として判定している

        https://github.com/rivo/uniseg/blob/master/width.go

    * このため2文字幅の環境ではいくつかの要素表示でずれる可能性があると思われる(2022年12月現在では未検証)

        https://github.com/rivo/tview/search?q=uniseg.StringWidth

    * それぞれの文字幅判定の検証結果は runewidth_eastasian_test.go のコードを参照のこと
