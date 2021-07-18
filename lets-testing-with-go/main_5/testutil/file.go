package testutil

import (
	"fmt"
	"io/ioutil"
	"syscall"
	"testing"
)

func TempFileNoCleanup(t *testing.T, content []byte) (name string, teardown func()) {
	file, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Error(err)
	}
	if err = ioutil.WriteFile(file.Name(), content, 0640); err != nil {
		t.Error(err)
	}
	return file.Name(), func() {
		fmt.Println("Call teardown()")
		syscall.Unlink(file.Name())
	}
}

func TempFileCleanup(t *testing.T, content []byte) string {
	file, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		fmt.Println("Call t.Cleanup()")
		syscall.Unlink(file.Name())
	})

	if err = ioutil.WriteFile(file.Name(), content, 0640); err != nil {
		t.Error(err)
	}
	return file.Name()
}
