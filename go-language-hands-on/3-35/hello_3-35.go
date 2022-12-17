package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type StData struct {
	msg string
	mux sync.Mutex
}

func main() {
	sd := StData{msg: "Start"}
	prmsg := func(nm string, n int) {
		fmt.Println(nm, sd.msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	main := func(n int) {
		const nm string = "*main"
		sd.mux.Lock() //☆
		for i := 0; i < 5; i++ {
			sd.msg += " m" + strconv.Itoa(i)
			prmsg(nm, n)
		}
		sd.mux.Unlock() //☆
	}

	hello := func(n int) {
		const nm string = "hello"
		sd.mux.Lock() //☆
		for i := 0; i < 5; i++ {
			sd.msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
		sd.mux.Unlock() //☆
	}

	go main(100)
	go hello(100)
	time.Sleep(5 * time.Second)
}
