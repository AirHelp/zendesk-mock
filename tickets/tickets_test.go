package tickets_test

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/api"
	"github.com/AirHelp/zendesk-mock/test"
	"github.com/AirHelp/zendesk-mock/tickets"
	"net/http/httptest"
	"testing"
)

func TestShow(t *testing.T) {
	var response = test.RecordGet("/api/v2/tickets/:id", "/api/v2/tickets/123", "", tickets.Show)

	test.IsExpectedToRespondWithCode(t, response, 200)
	test.IsExpectedToNotBeBlank(t, receivedTicket(response).Subject)
}

func TestCreate(t *testing.T) {
	var response = test.RecordPost("/api/v2/tickets", "/api/v2/tickets", body(), tickets.Create)

	test.IsExpectedToRespondWithCode(t, response, 201)
	test.IsExpectedToNotBeBlank(t, receivedTicket(response).Subject)
}

func TestUpdate(t *testing.T) {
	var response = test.RecordPut("/api/v2/tickets/:id", "/api/v2/tickets/1", body(), tickets.Update)

	test.IsExpectedToRespondWithCode(t, response, 200)
	test.IsExpectedToNotBeBlank(t, receivedTicket(response).Subject)
}

func body() string {
	return `{"ticket":{"subject":"Subject","comment":{"body":"Body"}}}`
}

func receivedTicket(response *httptest.ResponseRecorder) api.Ticket {
	var envelope api.TicketEnvelope
	dec := json.NewDecoder(response.Body)
	dec.Decode(&envelope)
	return envelope.Ticket
}
