package users_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AirHelp/zendesk-mock/users"
)

func TestUsersFindPath(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", users.UsersFindPath+"123", nil)
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
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"user":{"id":123,"name":"Name from Zendesk"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
