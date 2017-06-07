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

func TestDestroyMany(t *testing.T) {
	var response = test.RecordDelete(
		"/api/v2/group_memberships/destroy_many",
		"/api/v2/group_memberships/destroy_many?ids=1,2,3",
		"",
		group_memberships.DestroyMany)

	test.IsExpectedToRespondWithCode(t, response, 200)
}

func TestCreateMany(t *testing.T) {
	var response = test.RecordPost(
		"/api/v2/group_memberships/create_many",
		"/api/v2/group_memberships/create_many?ids=1,2,3",
		"",
		group_memberships.CreateMany)

	test.IsExpectedToRespondWithCode(t, response, 200)
}
