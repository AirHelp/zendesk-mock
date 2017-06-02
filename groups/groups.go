package groups

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/respond"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ApiUrl = "/api/v2/groups/"
)

type Response struct {
	Group Group `json:"group"`
}

// Group is used for a response
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type reqBody struct {
	Group reqGroup `json:"group"`
}

type reqGroup struct {
	Name string `json:"name"`
}

func Create(res http.ResponseWriter, req *http.Request) {
	if requestBody, err := RequestBody(req); err != nil {
		respond.WithJson(res, 400, nil, err)
	} else {
		RespondWithMock(res, int(time.Now().Unix()), requestBody.Group.Name)
	}
}

func Show(res http.ResponseWriter, req *http.Request) {
	if requestId, err := RequestId(req); err != nil {
		respond.WithJson(res, 404, nil, err)
	} else {
		RespondWithMock(res, requestId, "Group Name")
	}
}

func Update(res http.ResponseWriter, req *http.Request) {
	if requestBody, err := RequestBody(req); err != nil {
		respond.WithJson(res, 400, nil, err)
	} else if requestId, err := RequestId(req); err != nil {
		respond.WithJson(res, 404, nil, err)
	} else {
		RespondWithMock(res, requestId, requestBody.Group.Name)
	}
}

func RequestId(req *http.Request) (int, error) {
	return strconv.Atoi(strings.TrimPrefix(req.URL.Path, ApiUrl))
}

func RequestBody(req *http.Request) (reqBody, error) {
	var requestBody reqBody
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	return requestBody, err
}

func RespondWithMock(res http.ResponseWriter, id int, name string) {
	bytes, err := json.Marshal(Response{Group{ID: id, Name: name}})
	if err != nil {
		respond.WithJson(res, 500, nil, err)
	} else {
		respond.WithJson(res, 201, bytes, nil)
	}
}
