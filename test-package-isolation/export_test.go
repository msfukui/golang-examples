package mypkg

const ExportMaxValue = maxValue

func SetBaseUrl(s string) (resetFunc func()) {
	var tmp string
	tmp, baseURL = baseURL, s
	return func() {
		baseURL = tmp
	}
}

var ExportCounterReset = (*Counter).reset

func (c *Counter) ExportN() int {
	return c.n
}

func (c *Counter) ExportSetN(n int) {
	c.n = n
}
