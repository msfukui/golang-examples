package main

import (
	"encoding/json"
	"io/ioutil"
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
