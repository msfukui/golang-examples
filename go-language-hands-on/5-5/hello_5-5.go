package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// read text function.
	rt := func(f *os.File) {
		s, er := io.ReadAll(f)
		if er != nil {
			panic(er)
		}
		fmt.Println(string(s))
	}

	fn := "data.txt"

	f, er := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}

	// defer close
	defer f.Close()

	fmt.Println("<< start >>")
	rt(f)
	fmt.Println("<< end >>")
}
