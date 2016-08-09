package tickets

import (
	"encoding/json"
	"strconv"

	"github.com/go-martini/martini"
)

type response struct {
	Ticket ticket `json:"ticket"`
}

type ticket struct {
	Id      int    `json:"id"`
	Subject string `json:"subject"`
	Comment string `json:"comment"`
}

// New creates new ticket
func New(params martini.Params) (int, string) {
	comment := (params["comment"])
	response := response{ticket{12312, "Anything from Zendesk", comment}}
	bytes, err := json.Marshal(response)
	if err != nil {
		return 500, err.Error()
	}
	return 201, string(bytes)
}

// Find finds ticket
func Find(params martini.Params) (int, string) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return 404, err.Error()
	}
	response := response{ticket{id, "Anything from Zendesk", ""}}
	bytes, err := json.Marshal(response)
	if err != nil {
		return 500, err.Error()
	}
	return 200, string(bytes)
}
