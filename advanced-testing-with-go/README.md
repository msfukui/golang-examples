# advanced-testing-with-go

c.f. https://speakerdeck.com/mitchellh/advanced-testing-with-go

## よかったところ

* GOLDEN FILES のテストパターンを知ることができたのはよかった

* TEST HELPERS で後処理を関数で返して defer するパターンを知ることができた

* TestHelperProcess を使ったモック実装例はとてもかしこい！って思った

* `t.Parallel()` を使ったテストはやってしまいがちな気がする

    マシンを分けて明示的に実行テストを分ける、というのもちょっと違う気はする

## わからなかったところ

* faketime はみたところ取り扱い難しいし、timecop なものがるとうれしいのだろうか

    複数の実装があるみたい

    * https://github.com/Songmu/flextime

        https://songmu.jp/riji/entry/2020-01-19-flextime.html

    * https://github.com/agatan/timejump

        https://agtn.hatenablog.com/entry/2017/12/14/232124

    * https://github.com/bluele/go-timecop

    でもどのモジュールでも、グローバルで実行時に動的に差し替えることはできないので、素直に DI するのが良さそう

    c.f. https://twitter.com/r7kamura/status/390175256305876992
