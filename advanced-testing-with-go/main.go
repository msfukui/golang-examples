package main

import (
	"fmt"
	"time"
)

func Add(n int, m int) int {
	return n + m
}

func TimingDependentTests(n int) string {
	time.Sleep(time.Duration(n) * time.Second)
	return "done"
}

func ParallelTests(name string, value int) string {
	return fmt.Sprintf("%v, %v", name, value)
}
