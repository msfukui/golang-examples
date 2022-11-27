package main

import "net/http"

func main() {
	if err := http.ListenAndServe("", http.FileServer(http.Dir("."))); err != nil {
		panic(err)
	}
}
