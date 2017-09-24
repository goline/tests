package tests

import (
	"encoding/json"
	"io"
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
	return Call(http.MethodGet, uri, nil, output)
}

func Post(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPost, uri, strings.NewReader(body), output)
}

func Put(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPut, uri, strings.NewReader(body), output)
}

func Patch(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPatch, uri, strings.NewReader(body), output)
}

func Delete(uri string, output interface{}) *http.Response {
	return Call(http.MethodDelete, uri, nil, output)
}

func Call(method string, uri string, body io.Reader, output interface{}) *http.Response {
	r := httptest.NewRequest(method, uri, body)
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
