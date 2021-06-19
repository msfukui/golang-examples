package main

import "fmt"

func main() {
	msg := "Greetings\nfrom\nTurkey\n"

	count := countLines(msg)
	fmt.Println(count)
}

func countLines(msg string) int {
	var count int
	for i := 0; i < len(msg); i++ {
		if msg[i] == '\n' {
			count++
		}
	}
	return count
}
