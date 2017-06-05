package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type action func(res http.ResponseWriter, req *http.Request)

func RecordAction(t *testing.T, action string, url string, body string, a action) *httptest.ResponseRecorder {
	req, err := http.NewRequest(action, url, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a)

	handler.ServeHTTP(rr, req)
	return rr
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
