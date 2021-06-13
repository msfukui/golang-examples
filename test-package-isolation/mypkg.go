package mypkg

import "strings"

const maxValue = 100

var baseURL = "https://example.com/api/v2"

func Hoge() string {
	return "Hoge"
}

func DoSomething() int {
	return 99
}

func Client() string {
	return strings.Split(baseURL, ":")[0]
}

type Counter struct {
	n int
}

func (c *Counter) Count() {
	c.n++
}

func (c *Counter) reset() {
	c.n = 0
}

type response struct {
	Value string `json:"value"`
}

func (r *response) setResponse(s string) {
	r.Value = s
}

func (r *response) getResponse() string {
	return r.Value
}
