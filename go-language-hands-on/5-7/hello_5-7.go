package main

import (
	"fmt"
	"os"
)

func main() {
	es, er := os.ReadDir(".")
	if er != nil {
		panic(er)
	}
	for _, e := range es {
		f, _ := e.Info()
		fmt.Println(e.Name(), "(", f.Size(), ")")
	}
}
