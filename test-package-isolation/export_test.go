package mypkg

const ExportMaxValue = maxValue

func SetBaseUrl(s string) (resetFunc func()) {
	var tmp string
	tmp, baseURL = baseURL, s
	return func() {
		baseURL = tmp
	}
}
