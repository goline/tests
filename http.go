package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var handler http.Handler

func SetHandler(h http.Handler) {
	handler = h
}

func Get(uri string, output interface{}) *http.Response {
	return Call(http.MethodGet, uri, nil, output)
}

func Post(uri string, body interface{}, output interface{}) *http.Response {
	return Call(http.MethodPost, uri, body, output)
}

func Put(uri string, body interface{}, output interface{}) *http.Response {
	return Call(http.MethodPut, uri, body, output)
}

func Patch(uri string, body interface{}, output interface{}) *http.Response {
	return Call(http.MethodPatch, uri, body, output)
}

func Delete(uri string, output interface{}) *http.Response {
	return Call(http.MethodDelete, uri, nil, output)
}

func Call(method string, uri string, v interface{}, output interface{}) *http.Response {
	body, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	r := httptest.NewRequest(method, uri, bytes.NewReader(body))
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
