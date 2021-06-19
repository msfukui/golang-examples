package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(sayHi())
	fmt.Println(sayYoo())
}

// sayHi() returns the string "hi"
func sayHi() string {
	return "hi"
}

func sayYoo() string {
	return "yoo"
}
