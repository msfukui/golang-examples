package hello

import (
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	type args struct {
		in string
	}
	test := struct {
		name string
		args args
		in   string
		want string
	}{
		name: "sample", args: args{in: "type a number:"}, in: "1234", want: "1234",
	}

	funcDefer, err := mockStdin(t, test.in)
	if err != nil {
		t.Fatal(err)
	}

	defer funcDefer()

	if got := Input(test.args.in); got != test.want {
		t.Errorf("Input(\"%v\") = %v, want %v (mock %v)", test.args.in, got, test.want, test.in)
	}
}

// c.f. https://gist.github.com/KEINOS/76857bc6339515d7144e00f17adb1090
//
// mockStdin is a helper function that lets the test pretend dummyInput as os.Stdin.
// It will return a function for `defer` to clean up after the test.
func mockStdin(t *testing.T, dummyInput string) (funcDefer func(), err error) {
	t.Helper()

	oldOsStdin := os.Stdin

	tmpfile, err := os.CreateTemp(t.TempDir(), t.Name())
	if err != nil {
		return nil, err
	}

	content := []byte(dummyInput)

	if _, err := tmpfile.Write(content); err != nil {
		return nil, err
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, err
	}

	// Set stdin to the temp file
	os.Stdin = tmpfile

	return func() {
		// clean up
		os.Stdin = oldOsStdin
		os.Remove(tmpfile.Name())
	}, nil
}
