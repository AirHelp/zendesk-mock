package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-martini/martini"
)

const (
	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"
)

type action func(res http.ResponseWriter, req *http.Request, params martini.Params)

func RecordMethod(route, url, body, method string, handler action) *httptest.ResponseRecorder {
	m := martini.Classic()
	methods := map[string]func(string, ...martini.Handler) martini.Route{
		Get:    m.Get,
		Post:   m.Post,
		Put:    m.Put,
		Delete: m.Delete,
	}
	methods[method](route, handler)
	req, _ := http.NewRequest(method, url, strings.NewReader(body))
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
