package main

import (
	"testing"
)

func TestBar(t *testing.T) {
	result := Bar()
	if result != "foo" {
		t.Errorf("expecting foo, got %s", result)
	}
}

func TestQux(t *testing.T) {
	result := Qux("bar")
	if result != "foo" {
		t.Errorf("expecting foo, got %s", result)
	}

	result = Qux("qux")
	if result != "INVALID" {
		t.Errorf("expecting INVALID, got %s", result)
	}
}
