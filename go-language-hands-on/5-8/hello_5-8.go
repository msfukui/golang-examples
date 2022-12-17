package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	p := "https://golang.org"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	s, er := io.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}
	fmt.Println(string(s))
}
