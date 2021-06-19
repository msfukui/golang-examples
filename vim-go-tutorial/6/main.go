package main

import "fmt"

type Server struct {
	bar string
}

func main() {
	s := Server{bar: "Alper"}
	fmt.Println(s.bar) // print the server name
}

func name() string {
	return "Zeynep"
}
