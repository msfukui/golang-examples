package mypkg_test

import (
	"mypkg"
	"testing"
)

func TestHoge(t *testing.T) {
	if s := mypkg.Hoge(); s != "Hoge" {
		t.Error("want Hoge, got ", s)
	}
}

func TestMypkg(t *testing.T) {
	if s := mypkg.DoSomething(); s > mypkg.ExportMaxValue {
		t.Error("Error ", s)
	}
}
