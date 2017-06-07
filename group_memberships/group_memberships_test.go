package group_memberships_test

import (
	"github.com/AirHelp/zendesk-mock/group_memberships"
	"github.com/AirHelp/zendesk-mock/test"
	"testing"
)

func TestIndex(t *testing.T) {
	var response = test.RecordGet(
		"/api/v2/users/:user_id/group_memberships",
		"/api/v2/users/1/group_memberships",
		"",
		group_memberships.Index)

	test.IsExpectedToRespondWithCode(t, response, 200)
}
