package main

import (
	"fmt"
	"time"
)

func total(c chan int) {
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}

func main() {
	c := make(chan int)
	c <- 100
	go total(c)
	time.Sleep(100 * time.Millisecond)
}
