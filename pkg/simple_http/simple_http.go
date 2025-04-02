package simplehttp

import (
	"io"
	"net/http"
)

// map would be better, however we have only one param
func MakeRequest(Url string, Query string) ([]byte, error) {
	fullurl := Url + "?name=" + Query
	resp, err := http.Get(fullurl)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	return data, nil
}
