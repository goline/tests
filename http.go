package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var handler http.Handler

func SetHandler(h http.Handler) {
	handler = h
}

func Get(uri string, output interface{}) *http.Response {
	return Call(http.MethodGet, uri, "", output)
}

func Post(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPost, uri, body, output)
}

func Put(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPut, uri, body, output)
}

func Patch(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPatch, uri, body, output)
}

func Delete(uri string, output interface{}) *http.Response {
	return Call(http.MethodDelete, uri, "", output)
}

func Call(method string, uri string, body string, output interface{}) *http.Response {
	var buf *strings.Reader
	if body != "" {
		buf = strings.NewReader(body)
	}
	r := httptest.NewRequest(method, uri, buf)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	parseBody(w.Result(), output)
	return w.Result()
}

func parseBody(r *http.Response, output interface{}) {
	if output == nil {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, output)
	if err != nil {
		panic(err)
	}
}
