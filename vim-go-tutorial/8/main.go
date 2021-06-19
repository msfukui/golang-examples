package main

import "fmt"

type T struct{}

func (t *T) Read(p []byte) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

func (t *T) Write(p []byte) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

func (t *T) Close() error {
	panic("not implemented") // TODO: Implement
}

type B struct{}

func (b *B) String() string {
	panic("not implemented") // TODO: Implement
}

func main() {
	fmt.Println("vim-go")
}
