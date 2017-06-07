package group_memberships

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/mocks"
	"github.com/AirHelp/zendesk-mock/respond"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

func Index(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if userId, err := strconv.Atoi(params["user_id"]); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		respondWithCollection(res, 200, userId)
	}
}

func DestroyMany(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if req.URL.Query()["ids"] == nil {
		respond.Json(res, 400, nil, nil)
	} else {
		respondWithJobStatus(res, 200)
	}
}

func CreateMany(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if req.URL.Query()["ids"] == nil {
		respond.Json(res, 400, nil, nil)
	} else {
		respondWithJobStatus(res, 200)
	}
}

func respondWithCollection(res http.ResponseWriter, code int, userId int) {
	if bytes, err := json.Marshal(mocks.GroupMemberships(userId)); err != nil {
		respond.Json(res, 500, nil, err)
	} else {
		respond.Json(res, code, bytes, nil)
	}
}

func respondWithJobStatus(res http.ResponseWriter, code int) {
	if bytes, err := json.Marshal(mocks.JobStatus()); err != nil {
		respond.Json(res, 500, nil, err)
	} else {
		respond.Json(res, 200, bytes, nil)
	}
}
