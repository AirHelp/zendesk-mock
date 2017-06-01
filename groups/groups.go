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
	var input reqBody
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		respond.WithJson(res, 400, nil, err)
	} else {
		RespondWithMock(res, int(time.Now().Unix()), input.Group.Name)
	}
}

func Show(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(Id(req))
	if err != nil {
		respond.WithJson(res, 404, nil, err)
	} else {
		RespondWithMock(res, id, "Group Name")
	}
}

func Update(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(Id(req))
	if err != nil {
		respond.WithJson(res, 404, nil, err)
		return
	}
	var input reqBody
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		respond.WithJson(res, 400, nil, err)
	} else {
		RespondWithMock(res, id, input.Group.Name)
	}
}

func Id(req *http.Request) string {
	return strings.TrimPrefix(req.URL.Path, ApiUrl)
}

func RespondWithMock(res http.ResponseWriter, id int, name string) {
	responseBody := Response{Group{ID: id, Name: name}}
	bytes, err := json.Marshal(responseBody)
	if err != nil {
		respond.WithJson(res, 500, nil, err)
	} else {
		respond.WithJson(res, 201, bytes, nil)
	}
}
