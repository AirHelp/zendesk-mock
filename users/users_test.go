package users_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AirHelp/zendesk-mock/users"
)

func TestUsersFindPath(t *testing.T) {
	cases := []struct {
		uri, expected string
		code          int
	}{
		{users.UsersFindPath + "123", `{"user":{"id":123,"name":"Name from Zendesk"}}`, http.StatusOK},
		{users.UsersFindPath, "", 400},
	}
	for _, c := range cases {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("GET", c.uri, nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(users.Find)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if rr.Code != c.code {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, c.code)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != c.expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), c.expected)
		}
	}
}
