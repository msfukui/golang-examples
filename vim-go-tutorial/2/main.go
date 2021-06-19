package main

import "fmt"

type T struct {
	Foo string
}

func main() {
	t := T{
		Foo: "foo",
	}
	fmt.Printf("t = %+v\n", t)
}

func Bar() string {
	return "bar"
}

func BarFoo() string {
	return "bar_foo"
}
