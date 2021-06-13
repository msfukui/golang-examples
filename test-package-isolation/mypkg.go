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
