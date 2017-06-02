package groups_test

import (
	"github.com/AirHelp/zendesk-mock/groups"
	"github.com/AirHelp/zendesk-mock/test"
	"testing"
)

func TestGroupsShow(t *testing.T) {
	var response = test.RecordAction(t, "GET", "/api/v2/groups/123", "", groups.Show)

	test.IsExpectedToRespondWithCode(t, response, 200)
}

func TestGroupsCreate(t *testing.T) {
	var response = test.RecordAction(t, "POST", "/api/v2/groups", json(), groups.Create)

	test.IsExpectedToRespondWithCode(t, response, 201)
}

func TestGroupsUpdate(t *testing.T) {
	var response = test.RecordAction(t, "POST", "/api/v2/groups/1", json(), groups.Update)

	test.IsExpectedToRespondWithCode(t, response, 200)
}

func json() string {
	return `{"group":{"name":"Name"}}`
}
