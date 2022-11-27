package main

import "net/http"

func main() {
	if err := http.ListenAndServe("", http.NotFoundHandler()); err != nil {
		panic(err)
	}
}
