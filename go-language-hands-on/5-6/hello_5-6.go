package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read text function.
	rt := func(f *os.File) {
		r := bufio.NewReaderSize(f, 4096)
		for i := 1; true; i++ {
			s, _, er := r.ReadLine()
			if er != nil {
				break
			}
			fmt.Println(i, ":", string(s))
		}
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
