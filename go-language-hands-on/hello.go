package main

import (
	"fmt"
	"golang-example/go-language-hands-on/src/hello"
	"strconv"
)

func main() {
	x := hello.Input("type a price")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	p := float64(n)
	fmt.Println(int(p * 1.1))
}
