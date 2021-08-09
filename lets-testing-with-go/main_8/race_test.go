package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Fail by giving the test the `-race` option
func TestFoo(te *testing.T) {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Since(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
