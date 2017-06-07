package group_memberships_test

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/groups"
	"github.com/AirHelp/zendesk-mock/test"
	"net/http/httptest"
	"testing"
)

func TestGroupsShow(t *testing.T) {
	var response = test.RecordAction(t, "GET", "/api/v2/groups/123", "", groups.Show)

	test.IsExpectedToRespondWithCode(t, response, 200)
	test.IsExpectedToNotBeBlank(t, receivedGroup(response).Name)
}

func TestGroupsCreate(t *testing.T) {
	var response = test.RecordAction(t, "POST", "/api/v2/groups", body(), groups.Create)

	test.IsExpectedToRespondWithCode(t, response, 201)
	test.IsExpectedToNotBeBlank(t, receivedGroup(response).Name)
}

func TestGroupsUpdate(t *testing.T) {
	var response = test.RecordAction(t, "POST", "/api/v2/groups/1", body(), groups.Update)

	test.IsExpectedToRespondWithCode(t, response, 200)
	test.IsExpectedToNotBeBlank(t, receivedGroup(response).Name)
}

func body() string {
	return `{"group":{"name":"Name"}}`
}

func receivedGroup(response *httptest.ResponseRecorder) groups.Group {
	var envelope groups.Envelope
	dec := json.NewDecoder(response.Body)
	dec.Decode(&envelope)
	return envelope.Group
}
