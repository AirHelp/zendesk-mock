package tickets_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/AirHelp/zendesk-mock/tickets"
)

func TestTicketsNew(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/api/v2/tickets", strings.NewReader(`{"ticket":{"subject":"sample subject"}}`))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tickets.New)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 201 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 201)
	}

	// 	// Check the response body is what we expect.
	timestampID := int(time.Now().Unix())
	expected := fmt.Sprintf(`{"ticket":{"id":%v,"subject":"sample subject","comment":"","custom_fields":null}}`, timestampID)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestTicketsFindGet(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", tickets.TicketsFindURI+"123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tickets.Find)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"ticket":{"id":123,"subject":"Anything from Zendesk","comment":"","custom_fields":[{"id":23020926,"value":"ch_web"}]}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
