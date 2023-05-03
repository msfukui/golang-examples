package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

func main() {
	str := "つのだ☆HIRO"
	fmt.Printf("%s, %d\n", str, runewidth.StringWidth(str))

	border := "┌─────────┐"
	fmt.Printf("%s, %d\n", border, runewidth.StringWidth(border))
}
