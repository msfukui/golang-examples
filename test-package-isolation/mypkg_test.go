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

func TestClient(t *testing.T) {
	defer mypkg.SetBaseUrl("http://localhost:8080")()

	if s := mypkg.Client(); s != "http" {
		t.Error("Error ", s)
	}
}

func TestCounter(t *testing.T) {
	var c mypkg.Counter

	c.Count()
	c.Count()
	if mypkg.ExportCounterReset(&c); c.ExportN() != 0 {
		t.Error("Error ", c.ExportN())
	}

	c.Count()
	if c.ExportSetN(5); c.ExportN() != 5 {
		t.Error("Error ", c.ExportN())
	}
}

func TestResponse(t *testing.T) {
	var r mypkg.ExportResponse

	mypkg.ExportSetResponse(&r, "test")
	if mypkg.ExportGetResponse(&r) != "test" {
		t.Error("Error ", mypkg.ExportGetResponse(&r))
	}
}
