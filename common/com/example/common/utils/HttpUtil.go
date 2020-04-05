package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpPost(url string, args url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, args)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}
