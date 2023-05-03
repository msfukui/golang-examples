# bubbletea-sample

A sample codes in use `bubbletea`.

https://github.com/charmbracelet/bubbletea

* Windows Terminal で実行した場合に罫線の表示がずれる事象が発生

    * 結論: 実行前に RUNEWIDTH_EASTASIAN=0 を設定することで回避は可能

        example: $ RUNEWIDTH_EASTASIAN=0 go run examples/composable-views/main.go

    * 原因: East Asian Width における Ambiguous の扱いのずれの問題

    * Windows Terminal ではすべて半角文字扱いで強制される仕様の模様

        https://github.com/microsoft/terminal/pull/2928
        https://qiita.com/good_kobe/items/2da02b8d141984d1bc4d

    * bubble tea では描画用のライブラリとして lipgloss を明示的に使っている(内部で依存しているわけではない)

        https://github.com/charmbracelet/lipgloss

    * lipgloss は文字幅の判定時に go-runewidth を使っている(これは依存している)

        https://github.com/mattn/go-runewidth

    * go-runewidth は既定ではlocaleを見てcjkだと罫線については全角文字として判定している

    * 表示幅は全角として判定しているのに実際の表示は半角なので半分でしか表示されない事象が発生する

    * go-runewidth では環境変数 RUNEWIDTH_EASTASIAN に 1 以外を設定することで、強制的に半角文字として判定させることが可能

        c.f. https://noborus.github.io/blog/runewidth/
