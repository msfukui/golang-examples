package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			fmt.Printf("sample,  %v\n", i)
		case 1:
			break // meaningless
		case 2:
			fmt.Printf("example, %v\n", i)
		default:
			fmt.Printf("default, %v\n", i)
		}
	}
}
