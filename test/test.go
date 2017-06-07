package test

import (
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type action func(res http.ResponseWriter, req *http.Request, params martini.Params)

func RecordGet(route string, url string, body string, handler action) *httptest.ResponseRecorder {
	m := martini.Classic()
	m.Get(route, handler)
	req, _ := http.NewRequest("GET", url, strings.NewReader(body))
	return responseRecorder(m, req)
}

func RecordPost(route string, url string, body string, handler action) *httptest.ResponseRecorder {
	m := martini.Classic()
	m.Post(route, handler)
	req, _ := http.NewRequest("POST", url, strings.NewReader(body))
	return responseRecorder(m, req)
}

func RecordPut(route string, url string, body string, handler action) *httptest.ResponseRecorder {
	m := martini.Classic()
	m.Put(route, handler)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(body))
	return responseRecorder(m, req)
}

func RecordDelete(route string, url string, body string, handler action) *httptest.ResponseRecorder {
	m := martini.Classic()
	m.Delete(route, handler)
	req, _ := http.NewRequest("DELETE", url, strings.NewReader(body))
	return responseRecorder(m, req)
}

func responseRecorder(m *martini.ClassicMartini, req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	m.ServeHTTP(res, req)
	return res
}

func IsExpectedToRespondWithCode(t *testing.T, response *httptest.ResponseRecorder, code int) {
	if status := response.Code; status != code {
		t.Errorf("handler returned wrong status code: got %v want %v", status, code)
	}
}

func IsExpectedToNotBeBlank(t *testing.T, tested string) {
	if tested == "" {
		t.Errorf("Is expected to be present but is blank")
	}
}
