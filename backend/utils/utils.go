package utils

import (
	"strings"
)

/// create url based from params
func ConstructUrl(host string, endPoint string, params ...string) string {
	paramsV := "?"
	for _, v := range params {
		paramsV += v + "&"
	}
	return strings.Replace("https://"+host+"/"+endPoint+paramsV, " ", "%20", -1)
}
