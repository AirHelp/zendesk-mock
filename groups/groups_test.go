package groups_test

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/api"
	"github.com/AirHelp/zendesk-mock/groups"
	"github.com/AirHelp/zendesk-mock/test"
	"net/http/httptest"
	"testing"
)

func TestGroupsShow(t *testing.T) {
	var response = test.RecordGet("/api/v2/groups/:id", "/api/v2/groups/123", "", groups.Show)

	test.IsExpectedToRespondWithCode(t, response, 200)
	test.IsExpectedToNotBeBlank(t, receivedGroup(response).Name)
}

func TestGroupsCreate(t *testing.T) {
	var response = test.RecordPost("/api/v2/groups", "/api/v2/groups", body(), groups.Create)

	test.IsExpectedToRespondWithCode(t, response, 201)
	test.IsExpectedToNotBeBlank(t, receivedGroup(response).Name)
}

func TestGroupsUpdate(t *testing.T) {
	var response = test.RecordPut("/api/v2/groups/:id", "/api/v2/groups/1", body(), groups.Update)

	test.IsExpectedToRespondWithCode(t, response, 200)
	test.IsExpectedToNotBeBlank(t, receivedGroup(response).Name)
}

func body() string {
	return `{"group":{"name":"Name"}}`
}

func receivedGroup(response *httptest.ResponseRecorder) api.Group {
	var envelope api.GroupEnvelope
	dec := json.NewDecoder(response.Body)
	dec.Decode(&envelope)
	return envelope.Group
}
