package group_memberships

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/respond"
	"net/http"
	"strconv"
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

func Index(res http.ResponseWriter, req *http.Request) {
	if requestUserId, err := requestUserId(req); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		RespondWithCollectionOfMocks(res, 200, requestUserId)
	}
}

func RespondWithCollectionOfMocks(res http.ResponseWriter, code int, userId int) {
	bytes, err := json.Marshal(mockedCollection(userId))
	if err != nil {
		respond.Json(res, 500, nil, err)
	} else {
		respond.Json(res, code, bytes, nil)
	}
}

func mockedCollection(userId int) CollectionEnvelope {
	collection := []GroupMembership{
		GroupMembership{Id: 1, Default: false, UserId: userId, GroupId: 1 },
		GroupMembership{Id: 1, Default: false, UserId: userId, GroupId: 2 }}
	return CollectionEnvelope{collection}
}

func requestUserId(req *http.Request) (int, error) {
	// return strconv.Atoi(strings.TrimPrefix(req.URL.Path, ApiUrl))
	return strconv.Atoi("1")
}
