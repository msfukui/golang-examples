// +build darwin

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type foo struct {
	Message    string
	Ports      []int
	ServerName string
}

func newFoo() (*foo, error) {
	return &foo{
		Message:    "foo loves bar",
		Ports:      []int{80},
		ServerName: "Foo",
	}, nil
}

func main() {
	fmt.Println(strings.ToUpper("gopher"))
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)

	res, err := newFoo()
	if err != nil {
		panic(err)
	}

	out, _ := json.Marshal(res)
	fmt.Println(string(out))
}

// bar returns a the string "foo" even though it's named as "bar". It's an
// example to be used with vim-go's tutorial to show the 'if' and 'af' text
// objects.
func Bar() string {
	fmt.Println("calling bar")

	foo := func() string {
		return "foo"
	}

	return foo()
}

func Foo() string {
	return "foo"
}

func Qux(v string) string {
	if v == "foo" {
		return Foo()
	}

	if v == "bar" {
		return Bar()
	}

	return "INVALID"
}
