package main_5

import (
	"example/main_5/testutil"
	"fmt"
	"os"
	"testing"
)

func f() {
	fmt.Println("なんらかの処理")
}

func Test_f(t *testing.T) {
	f()
}

func TestMain(m *testing.M) {
	fmt.Println("前処理")
	status := m.Run()
	fmt.Println("後処理")
	os.Exit(status)
}

func TestTempFileNoCleanup(t *testing.T) {
	file, delete := testutil.TempFileNoCleanup(t, nil)
	fmt.Println("TempFileNoCleanup(): " + file)
	defer delete()
}

func TestTempFileCleanup(t *testing.T) {
	file := testutil.TempFileCleanup(t, nil)
	fmt.Println("TempFileNoCleanup(): " + file)
}
