package utils

import (
	"errors"
	"net/http"
	"strconv"
)

//检查是否有version
func CheckVersionValue(query UrlValue, writer http.ResponseWriter) (int, error) {
	var versionString = query.Get("version")
	if versionString == "" {
		return 0, errors.New("version 不能空！")
	}
	var version, _ = strconv.Atoi(versionString)
	return version, nil
}
