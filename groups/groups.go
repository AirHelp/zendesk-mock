package groups

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/api"
	"github.com/AirHelp/zendesk-mock/mocks"
	"github.com/AirHelp/zendesk-mock/respond"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

func Create(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if requestBody, err := RequestBody(req); err != nil {
		respond.Json(res, 400, nil, err)
	} else {
		RespondWithMock(res, 201, mocks.Id(), requestBody.Group.Name)
	}
}

func Show(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if id, err := strconv.Atoi(params["id"]); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		RespondWithMock(res, 200, id, "Group Name")
	}
}

func Update(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if requestBody, err := RequestBody(req); err != nil {
		respond.Json(res, 400, nil, err)
	} else if id, err := strconv.Atoi(params["id"]); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		RespondWithMock(res, 200, id, requestBody.Group.Name)
	}
}

func RequestBody(req *http.Request) (api.GroupEnvelope, error) {
	var envelope api.GroupEnvelope
	err := json.NewDecoder(req.Body).Decode(&envelope)
	return envelope, err
}

func RespondWithMock(res http.ResponseWriter, code int, id int, name string) {
	bytes, err := json.Marshal(mocks.Group(id, name))
	if err != nil {
		respond.Json(res, 500, nil, err)
	} else {
		respond.Json(res, code, bytes, nil)
	}
}
