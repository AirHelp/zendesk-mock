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

type Envelope struct {
	Group Group `json:"group"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Create(res http.ResponseWriter, req *http.Request) {
	if requestBody, err := RequestBody(req); err != nil {
		respond.Json(res, 400, nil, err)
	} else {
		RespondWithMock(res, 201, int(time.Now().Unix()), requestBody.Group.Name)
	}
}

func Show(res http.ResponseWriter, req *http.Request) {
	if requestId, err := RequestId(req); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		RespondWithMock(res, 200, requestId, "Group Name")
	}
}

func Update(res http.ResponseWriter, req *http.Request) {
	if requestBody, err := RequestBody(req); err != nil {
		respond.Json(res, 400, nil, err)
	} else if requestId, err := RequestId(req); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		RespondWithMock(res, 200, requestId, requestBody.Group.Name)
	}
}

func RequestId(req *http.Request) (int, error) {
	return strconv.Atoi(strings.TrimPrefix(req.URL.Path, ApiUrl))
}

func RequestBody(req *http.Request) (Envelope, error) {
	var envelope Envelope
	err := json.NewDecoder(req.Body).Decode(&envelope)
	return envelope, err
}

func RespondWithMock(res http.ResponseWriter, code int, id int, name string) {
	bytes, err := json.Marshal(Envelope{Group{Id: id, Name: name}})
	if err != nil {
		respond.Json(res, 500, nil, err)
	} else {
		respond.Json(res, code, bytes, nil)
	}
}
