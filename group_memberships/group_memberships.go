package group_memberships

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/respond"
	"net/http"
	"strconv"
	"github.com/AirHelp/zendesk-mock/mocks"
	"github.com/go-martini/martini"
)

const (
	ApiUrl = "/api/v2/groups/"
)

type CollectionEnvelope struct {
	GroupMemberships []GroupMembership `json:"group_memberships"`
}

type GroupMembership struct {
	Id   int    	`json:"id"`
	Default bool 	`json:"default"`
	UserId int    `json:"user_id"`
	GroupId int   `json:"group_id"`
}

func Index(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if userId, err := strconv.Atoi(params["user_id"]); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		respondWithMockedCollection(res, 200, userId)
	}
}

func respondWithMockedCollection(res http.ResponseWriter, code int, userId int) {
	if bytes, err := json.Marshal(mockedCollection(userId)); err != nil {
		respond.Json(res, 500, nil, err)
	} else {
		respond.Json(res, code, bytes, nil)
	}
}

func mockedCollection(userId int) CollectionEnvelope {
	collection := []GroupMembership{
		GroupMembership{Id: mocks.Id(), Default: false, UserId: userId, GroupId: mocks.Id()},
		GroupMembership{Id: mocks.Id() + 1, Default: false, UserId: userId, GroupId: mocks.Id() + 1}}
	return CollectionEnvelope{collection}
}
