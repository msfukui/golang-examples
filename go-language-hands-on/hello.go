package main

import (
	"fmt"
	"golang-example/go-language-hands-on/src/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
