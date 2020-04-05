package core_util

import "strings"

func CutJsonString(result string) string {
	if result == "" || len(result) <= 2 {
		return ""
	}
	result = string([]byte(result)[1:(len(result) - 1)])
	result = strings.Replace(result, `\"`, `"`, -1)
	return result
}
