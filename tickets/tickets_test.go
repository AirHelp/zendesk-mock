package tickets_test

import (
	"encoding/json"
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

	// Check the response body is what we expect.
	timestampID := int(time.Now().Unix())
	var ticketResponse tickets.Response
	dec := json.NewDecoder(rr.Body)
	err = dec.Decode(&ticketResponse)
	if err != nil {
		t.Errorf("Unable to unmarshal response: %v", rr.Body.String())
	}
	if ticketResponse.Ticket.ID < timestampID {
		t.Errorf("Wrong ID in response: %v", ticketResponse.Ticket.ID)
	}
}

func TestTicketsFindGet(t *testing.T) {
	cases := []struct {
		uri, expected string
		code          int
	}{
		{
			tickets.TicketsFindURI + "123",
			`{"ticket":{"id":123,"subject":"Anything from Zendesk","comment":"","custom_fields":[{"id":28367069,"value":"GF4534"},{"id":28518725,"value":"delayed"},{"id":28367089,"value":"600"},{"id":28518845,"value":"600.0"},{"id":28518825,"value":"600"},{"id":28367689,"value":"600.0"},{"id":28367989,"value":"en"},{"id":28368009,"value":"ch_web"},{"id":28518885,"value":"John Doe"},{"id":28518905,"value":"John F. Kennedy International Airport, New York (JFK)"},{"id":28368029,"value":"Tegel Airport, Berlin (TXL)"},{"id":28518945,"value":"LOT - Polish Airlines (LO)"},{"id":28368069,"value":"554"},{"id":28368089,"value":"2017-03-13"},{"id":28518965,"value":"LO"},{"id":29960329,"value":""}]}}`,
			200,
		},
		{
			tickets.TicketsFindURI,
			"",
			404,
		},
	}
	for _, c := range cases {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("GET", c.uri, nil)
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
