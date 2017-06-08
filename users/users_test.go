package users_test

import (
	"github.com/AirHelp/zendesk-mock/test"
	"github.com/AirHelp/zendesk-mock/users"
	"testing"
)

func TestUsersShow(t *testing.T) {
	var response = test.RecordMethod("/api/v2/users/:id", "/api/v2/users/123", "", test.Get, users.Show)

	test.IsExpectedToRespondWithCode(t, response, 200)
}

func TestUsersShowNotExists(t *testing.T) {
	var response = test.RecordMethod("/api/v2/users/:id", "/api/v2/users/wer", "", test.Get, users.Show)

	test.IsExpectedToRespondWithCode(t, response, 404)
}
