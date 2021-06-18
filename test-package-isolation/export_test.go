package mypkg

const ExportMaxValue = maxValue

func SetBaseUrl(s string) (resetFunc func()) {
	var tmp string
	tmp, mybaseURL = mybaseURL, s
	return func() {
		mybaseURL = tmp
	}
}

var ExportCounterReset = (*Counter).reset

func (c *Counter) ExportN() int {
	return c.n
}

func (c *Counter) ExportSetN(n int) {
	c.n = n
}

type ExportResponse = response

var ExportSetMyResponse = (*response).setResponse
var ExportGetMyResponse = (*response).getResponse

func SetBaseURL(s string) (resetFunc func()) {
	var tmp string
	tmp, baseURL = baseURL, s
	return func() {
		baseURL = tmp
	}
}

type ExportGetResponse = getResponse
