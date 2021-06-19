package main

import (
	"example/example"
	"fmt"
)

func main() {
	Hello(example.GopherCon)
	Hello(example.Kenya)
}

func Hello(fn func() string) {
	fmt.Println("Hello " + fn())
}
