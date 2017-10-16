package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"fmt"
)

var handler http.Handler

func SetHandler(h http.Handler) {
	handler = h
}

func Get(uri string, output interface{}) *http.Response {
	return Call(http.MethodGet, uri, "", nil, output)
}

func Post(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPost, uri, body, nil, output)
}

func Put(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPut, uri, body, nil, output)
}

func Patch(uri string, body string, output interface{}) *http.Response {
	return Call(http.MethodPatch, uri, body, nil, output)
}

func Delete(uri string, output interface{}) *http.Response {
	return Call(http.MethodDelete, uri, "", nil, output)
}

func Call(method string, uri string, body string, header http.Header, output interface{}) *http.Response {
	r := httptest.NewRequest(method, uri, strings.NewReader(body))
	if header != nil {
		r.Header = header
	}
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

	if ShowResponse {
		fmt.Printf("%s\n", body)
	}

	err = json.Unmarshal(body, output)
	if err != nil {
		panic(err)
	}
}
