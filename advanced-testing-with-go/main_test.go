package main

// Sample code - Advanced Testing with Go
// https://speakerdeck.com/mitchellh/advanced-testing-with-go

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
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

// NETWORKING - 37
// Error checking omitted for brevity
// net.Conn をモックにする必要はないよ、
// モックではなく実際に通信を作ってテストしよう、
// と書かれている。
func TestConn(t *testing.T) {
	ln, _ := net.Listen("tcp", "127.0.0.1:0")

	go func() {
		defer ln.Close()
		ln.Accept()
	}()

	net.Dial("tcp", ln.Addr().String())
}

// SUBPROCESSING: REAL - 45
// ない場合はモックを使わずに明示的にテストをスキップする
// 副作用に注意
var testHasGit bool

func init() {
	if _, err := exec.LookPath("git"); err == nil {
		testHasGit = true
	}
}

func TestGitGeter(t *testing.T) {
	if !testHasGit {
		t.Log("git not found, skipping")
		t.Skip()
	}
}

// SUBPROCESSING: MOCK - 47, 48, 49
// See: src/os/exec/exec/exec_test.go: helperCommandContext()
// 内部で自身の TestHelperProcess() のみを引数付きで呼び出す様に
// exec.Cmd に差し込んでいる
func helperProcess(s ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cs = append(cs, s...)
	env := []string{
		"GO_WANT_HELPER_PROCESS=1",
	}

	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = append(env, os.Environ()...)
	return cmd
}

// see: src/os/exec/exec/exec_test.go: TestHelperProcess()
// TestHelperProcess() にはそれぞれのコマンドに対応したモックの実装を書く
func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}

	cmd, args := args[0], args[1:]
	switch cmd {
	case "echo":
		iargs := []interface{}{}
		for _, s := range args {
			iargs = append(iargs, s)
		}
		fmt.Println(iargs...)
	case "cat":
		if len(args) == 0 {
			io.Copy(os.Stdout, os.Stdin)
			return
		}
		exit := 0
		for _, fn := range args {
			f, err := os.Open(fn)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				exit = 2
			} else {
				defer f.Close()
				io.Copy(os.Stdout, f)
			}
		}
		os.Exit(exit)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}
}

// Use helperProcess(): echo
// モックを使ったテストの実例: echo の場合
func TestEcho(t *testing.T) {
	bs, err := helperProcess("echo", "foo bar", "baz").Output()
	if err != nil {
		t.Errorf("echo: %v", err)
	}
	if g, e := string(bs), "foo bar baz\n"; g != e {
		t.Errorf("echo: want %q, got %q", e, g)
	}
}

// Use helperProcess(): cat
// モックを使ったテストの実例: cat の場合
func TestCatStdin(t *testing.T) {
	// Cat, testing stdin and stdout.
	input := "Input string\nLine 2"
	p := helperProcess("cat")
	p.Stdin = strings.NewReader(input)
	bs, err := p.Output()
	if err != nil {
		t.Errorf("cat: %v", err)
	}
	s := string(bs)
	if s != input {
		t.Errorf("cat: want %q, got %q", input, s)
	}
}

// TIMING DEPENDENT TESTS - 64
// 記事では fake time は使うな、と書かれていたが、
// 別記事から持ってきてタイミングテストのサンプルを書いてみた
// 実行時にfaketime を使う場合は `-tags faketime` または `--tags=faketime` を付与する
// `net` などネットワーク関連のモジュールを使っているとハングするらしいので注意
// c.f. https://qiita.com/hogedigo/items/c2b6281961c5e21c4907
func TestTimingDependentTests(t *testing.T) {
	n := 1
	start := time.Now()
	result := TimingDependentTests(n)
	if expected := "done"; result != expected {
		t.Fatalf("TimingDependantTests(): want %v, got %v", expected, result)
	}
	duration := time.Since(start)
	delta := duration - time.Duration(n)*time.Second
	if delta < 0 {
		delta *= -1
	}
	if delta > time.Second {
		t.Fatalf("TimingDependantTests(): Illigal duration: %v", duration)
	}
}

// PARALLELIZATION - 66
// 記事では並列化はするな、と書かれていたが、
// 別記事から持ってきて並列化テストのサンプルを書いてみた
// 記事内だとハマるケースみたいだけど、テスト対象を関数にした場合問題なく動作した
// c.f. https://zenn.dev/ucwork/articles/cd26d933978080
func TestParallelTests(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{name: "test1", value: 1, expected: "test1, 1"},
		{name: "test2", value: 2, expected: "test2, 2"},
		{name: "test3", value: 3, expected: "test3, 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParallelTests(tt.name, tt.value)
			fmt.Println(result)
			if result != tt.expected {
				t.Fatalf("ParallelTests(): want %v, got %v", tt.expected, result)
			}
		})
	}
}
