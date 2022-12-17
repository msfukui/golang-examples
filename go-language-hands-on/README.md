# go-language-hands-on

「Go言語ハンズオン」 (ISBN978-4-7980-6399-7) を進めるためのディレクトリです.

## 最初に

以下をコマンドラインで実行して,パッケージ名を検索できるようにしておきます.

```
$ go mod init golang-example/go-language-hands-on
$ go get github.com/PuerkitoBio/goquery # <- for chapter 5.2
$ go get github.com/mattn/go-sqlite3 # <- for chapter 5.3
```

5.3 で使われる SQLite3 のデータは以下の様に用意します。

```
$ sqlite3 data.sqlite3
SQLite version 3.31.1 2020-01-27 19:55:54
Enter ".help" for usage hints.
sqlite> create table "mydata" (
   ...> "id" INTEGER PRIMARY KEY AUTOINCREMENT,
   ...> "name" TEXT NOT NULL,
   ...> "mail" TEXT,
   ...> "age" INTEGER
   ...> );
sqlite> INSERT INTO "mydata" VALUES(1,'Taro','taro@yamada',39);
sqlite> INSERT INTO "mydata" VALUES(2,'Hanako','hanako@flower',28);
sqlite> INSERT INTO "mydata" VALUES(3,'Sachiko','sachiko@happy',17);
sqlite> INSERT INTO "mydata" VALUES(4,'Jiro','jiro@change',6);
sqlite> select * from mydata ;
1|Taro|taro@yamada|39
2|Hanako|hanako@flower|28
3|Sachiko|sachiko@happy|17
4|Jiro|jiro@change|6
sqlite> .quit
```
