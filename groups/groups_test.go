package groups_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/AirHelp/zendesk-mock/api"
	"github.com/AirHelp/zendesk-mock/groups"
	"github.com/AirHelp/zendesk-mock/test"
)

func TestGroupsShow(t *testing.T) {
	var response = test.RecordMethod("/api/v2/groups/:id", "/api/v2/groups/123", "", test.Get, groups.Show)

	test.IsExpectedToRespondWithCode(t, response, 200)
	test.IsExpectedToNotBeBlank(t, receivedGroup(response).Name)
}

func TestGroupsCreate(t *testing.T) {
	var response = test.RecordMethod("/api/v2/groups", "/api/v2/groups", body(), test.Post, groups.Create)

	test.IsExpectedToRespondWithCode(t, response, 201)
	test.IsExpectedToNotBeBlank(t, receivedGroup(response).Name)
}

func TestGroupsUpdate(t *testing.T) {
	var response = test.RecordMethod("/api/v2/groups/:id", "/api/v2/groups/1", body(), test.Put, groups.Update)

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
