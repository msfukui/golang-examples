package main

import (
	"net/url"
	"testing"
)

// For test Helper function
func mustUrlParse(t *testing.T, s string) *url.URL {
	t.Helper()
	u, err := url.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	return u
}

func TestX(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal1", args{"http://example.com/"}, 0},
		{"normal2", args{"%zzzzz" /* error occured */}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := mustUrlParse(t, tt.args.str)
			if url.String() == "" {
				t.Error("url is empty!")
			}
		})
	}
}
