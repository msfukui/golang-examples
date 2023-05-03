package main

import (
	"testing"

	"github.com/mattn/go-runewidth"
	"github.com/rivo/uniseg"
)

func TestUnisegStringWidth(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "multi byte string, Ambiguous is 1", args: args{text: "△1▽23"}, want: 5},
		{name: "multi byte string, Ambiguous is 2", args: args{text: "あ1い23"}, want: 7},
		{name: "multi byte string, Ambiguous is 1, 2", args: args{text: "つのだ☆HIRO"}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniseg.StringWidth(tt.args.text); got != tt.want {
				t.Errorf("uniseg.StringWidth(%v) = %v, want %v", tt.args.text, got, tt.want)
			}
		})
	}
}

func TestRuneWidthStringWidth(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "multi byte string, Ambiguous is 1", args: args{text: "△1▽23"}, want: 5}, // RUNEWIDTH_EASTASIAN=1, want: 7
		{name: "multi byte string, Ambiguous is 2", args: args{text: "あ1い23"}, want: 7},
		{name: "multi byte string, Ambiguous is 1, 2", args: args{text: "つのだ☆HIRO"}, want: 11}, // RUNEWIDTH_EASTASIAN=1, want: 12
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runewidth.StringWidth(tt.args.text); got != tt.want {
				t.Errorf("runewidth.StringWidth(%v) = %v, want %v", tt.args.text, got, tt.want)
			}
		})
	}
}
