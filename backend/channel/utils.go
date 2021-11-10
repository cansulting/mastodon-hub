package channel

func getUrlFrom(host string, endPoint string, params ...string) string {
	paramsV := "?"
	for _, v := range params {
		paramsV += v + "&"
	}
	return "http://" + host + "/" + endPoint + paramsV
}
