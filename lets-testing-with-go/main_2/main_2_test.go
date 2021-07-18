package main_2

import "testing"

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "fail", args: args{a: 1, b: 2}, want: 30},
		{name: "normal", args: args{a: 1, b: 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			// t.Fatalf でテストが失敗した場合でも defer の処理は呼び出される
			defer t.Log("defer!")

			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Fatalf("add() = %v, want %v", got, tt.want)
			}

			// t.Fatalf でテストが失敗した場合は以下は呼び出されない
			t.Log("after add() ...")
		})
	}
}
