package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	p := "https://golang.org"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	// Depricated the goquery.NewDocument() from goquery v1.4.0.
	// c.f. https://github.com/PuerkitoBio/goquery/issues/173
	// doc, er := goquery.NewDocument(p)
	doc, er := goquery.NewDocumentFromReader(re.Body)
	if er != nil {
		panic(er)
	}

	doc.Find("a").Each(func(n int, sel *goquery.Selection) {
		lk, _ := sel.Attr("href")
		println(n, sel.Text(), "(", lk, ")")
	})
}
