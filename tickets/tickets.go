package tickets

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
	if requestBody, err := requestBody(req); err != nil {
		respond.Json(res, 400, nil, err)
	} else {
		respondWithMock(res, 201, mocks.Id(), requestBody.Ticket.Subject)
	}
}

func Show(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if id, err := strconv.Atoi(params["id"]); err != nil {
		respond.Json(res, 404, nil, err)
	} else {
		respondWithMock(res, 200, id, "Ticket Subject")
	}
}

func Update(res http.ResponseWriter, req *http.Request, params martini.Params) {
	if id, err := strconv.Atoi(params["id"]); err != nil {
		respond.Json(res, 404, nil, err)
	} else if requestBody, err := requestBody(req); err != nil {
		respond.Json(res, 400, nil, err)
	} else {
		respondWithMock(res, 200, id, requestBody.Ticket.Subject)
	}
}

func requestBody(req *http.Request) (api.CreateTicketEnvelope, error) {
	var envelope api.CreateTicketEnvelope
	err := json.NewDecoder(req.Body).Decode(&envelope)
	return envelope, err
}

func respondWithMock(res http.ResponseWriter, code int, id int, subject string) {
	bytes, err := json.Marshal(mocks.Ticket(id, subject))
	if err != nil {
		respond.Json(res, 500, nil, err)
	} else {
		respond.Json(res, code, bytes, nil)
	}
}
