package users

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/mocks"
	"github.com/AirHelp/zendesk-mock/respond"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

func Show(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if id, err := strconv.Atoi(params["id"]); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		respondWithMock(res, 200, id, "User Name")
	}
}

func respondWithMock(res http.ResponseWriter, code int, id int, name string) {
	if bytes, err := json.Marshal(mocks.User(id, name)); err != nil {
		respond.Json(res, 500, nil, err)
	} else {
		respond.Json(res, code, bytes, nil)
	}
}
