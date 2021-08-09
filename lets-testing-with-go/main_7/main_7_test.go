package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// from APIサーバにアクセスするテストをしたい
func TestX(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("GET %s: expected status code = %d; got %d", ts.URL, 200, res.StatusCode)
	}
	if string(got) != "Hello, client\n" {
		t.Errorf("expected body %v; got %v", "Hello, client", string(got))
	}
}

// from APIサーバのハンドラのテストをしたい
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}

func TestHelloHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/dummy", nil)
	w := httptest.NewRecorder()

	helloHandler(w, r)

	resp := w.Result()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("cannot read test response: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("got = %d, want = 200", resp.StatusCode)
	}

	if string(body) != "hello world!" {
		t.Errorf("got = %s, want = hello world!", body)
	}
}
