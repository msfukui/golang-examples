package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// TABLE DRIVEN TESTS - 13
func TestAdd(t *testing.T) {
	cases := []struct{ A, B, Expected int }{
		{1, 1, 2},
		{1, -1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}

	for _, tc := range cases {
		actual := Add(tc.A, tc.B)
		if actual != tc.Expected {
			t.Errorf(
				"Add(%d, %d) = %d, expected %d",
				tc.A, tc.B, actual, tc.Expected,
			)
		}
	}
}

// TABLE DRIVEN TESTS - 15
func TestAdd2(t *testing.T) {
	cases := map[string]struct{ A, B, Expected int }{
		"foo": {1, 1, 2},
		"bar": {1, -1, 0},
	}

	for k, tc := range cases {
		actual := Add(tc.A, tc.B)
		if actual != tc.Expected {
			t.Errorf(
				"%s: Add(%d, %d) = %d, expected %d",
				k, tc.A, tc.B, actual, tc.Expected,
			)
		}
	}
}

type Case struct {
	Label    string
	A        int
	B        int
	Expected int
}

// TEST FIXTURES - 17
func TestAdd3(t *testing.T) {
	const N = 2
	cases := make([]Case, N)

	for i := 0; i < N; i++ {
		data := filepath.Join("test-fixtures", "add_data_"+strconv.Itoa(i+1)+".json")

		raw, err := ioutil.ReadFile(data)
		if err != nil {
			t.Errorf("do not read a fixture file %s", data)
		}

		err = json.Unmarshal(raw, &cases[i])
		if err != nil {
			t.Error("do not unmarshal json data")
		}
	}

	for _, v := range cases {
		actual := Add(v.A, v.B)
		if actual != v.Expected {
			t.Errorf(
				"%s: Add(%d, %d) = %d, expected %d",
				v.Label, v.A, v.B, actual, v.Expected,
			)
		}
	}
}

// GOLDEN FILES - 20
// テスト結果を update フラグがついていれば、ファイルに保存できる様にした上で、
// 普段は、その保存しておいたファイルから期待値を読み取ってテストを実行する。
// `go test -update` でフラグを立てて実行できる。
// 利用において以下3点のコメントが記載されている。
// * 手動でハードコーディングせずに複雑な出力をテストする
// * 生成されたデータを目で見て、正しい場合は、データをコミットする
// * 複雑な構造をテストするためのスケーラブルな方法(String()を実装)
//
// テストの期待値を手で書かず容易に管理・実現できるメリットがありそう。
// ただしデータの保管の実装は少し煩雑になるため、byteデータの変換などでヘルパーの実装が必要そう。
// このサンプルではbyteで保管しているが、実際には、目視で確認するためにtextで保管した方がよさそう。
//
var update = flag.Bool("update", false, "update golden files")

func TestAdd4(t *testing.T) {
	cases := map[string]struct{ A, B int }{
		"foo": {1, 1},
		"bar": {1, -1},
	}

	for k, tc := range cases {
		actual := Add(tc.A, tc.B)
		actual64 := int64(actual)
		byteActual := make([]byte, binary.MaxVarintLen64)
		binary.PutVarint(byteActual, actual64)

		golden := filepath.Join("test-fixtures", k+".golden")
		if *update {
			ioutil.WriteFile(golden, byteActual, 0640)
		}

		byteExpected, _ := ioutil.ReadFile(golden)
		if !bytes.Equal(byteActual, byteExpected) {
			t.Errorf(
				"%s: Add(%d, %d) = %v, expected %v",
				k, tc.A, tc.B, byteActual, byteExpected,
			)
		}
	}
}

// TEST HELPERS - 29
// 後始末用の関数を返却して defer で終了時に遅延実行する実装例。
func testTempFile(t *testing.T) (string, func()) {
	tf, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return tf.Name(), func() { os.Remove(tf.Name()) }
}

func TestThing(t *testing.T) {
	tf, tfClose := testTempFile(t)
	defer tfClose()

	if tf == "" {
		t.Errorf("tf.Name(): %s", tf)
	}
}

// TEST HELPERS - 30
// クロージャーを使った実装例。
func testChdir(t *testing.T, dir string) func() {
	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("err: %s", err)
	}

	return func() { os.Chdir(old) }
}

func TestThing2(t *testing.T) {
	defer testChdir(t, "/tmp")()

	if actual := Add(1, 1); actual != 2 {
		t.Errorf("actual: %d", actual)
	}
}
