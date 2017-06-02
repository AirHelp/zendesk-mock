package groups_test

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/groups"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGroupsCreate(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/api/v2/groups", strings.NewReader(`{"group":{"name":"Name"}}`))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(groups.Create)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 201 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 201)
	}

	// Check the response body is what we expect.
	timestampID := int(time.Now().Unix())
	var response groups.Envelope
	dec := json.NewDecoder(rr.Body)
	err = dec.Decode(&response)
	if err != nil {
		t.Errorf("Unable to unmarshal response: %v", rr.Body.String())
	}
	if response.Group.Id < timestampID {
		t.Errorf("Wrong ID in response: %v", response.Group.Id)
	}
}
