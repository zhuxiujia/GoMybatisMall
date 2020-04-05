package utils

import (
	"encoding/json"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func JsonFunc(argFunc func(writer http.ResponseWriter, request *http.Request) interface{}) func(writer http.ResponseWriter, request *http.Request) {
	var f = func(writer http.ResponseWriter, request *http.Request) {
		var r = argFunc(writer, request)
		if r != nil {
			var rv = reflect.TypeOf(r)
			if strings.Contains(rv.String(), "error") || strings.Contains(rv.String(), "Error") {
				var msg = r.(error).Error()
				var v = vo.ResultVO{}.NewError(-1, msg)
				var b, _ = json.Marshal(v)
				writer.Header().Set("Content-type", "application/json")
				writer.Write(b)
			} else {
				var b, _ = json.Marshal(r)
				writer.Header().Set("Content-type", "application/json")
				writer.Write(b)
			}
		}
	}

	return f
}

type UrlValue struct {
	*url.Values
}

func (it UrlValue) GetString(k string) string {
	return it.Values.Get(k)
}
func (it *UrlValue) GetInt(k string, def int) int {
	var v = it.GetString(k)
	var r, e = strconv.Atoi(v)
	if e != nil {
		return def
	}
	return r
}
func (it *UrlValue) GetFloat(k string, def float64) float64 {
	var v = it.GetString(k)
	var r, e = strconv.ParseFloat(v, 64)
	if e != nil {
		return def
	}
	return r
}

func ParserUrlValues(request *http.Request) UrlValue {
	request.ParseForm()
	return UrlValue{Values: &request.Form}
}
