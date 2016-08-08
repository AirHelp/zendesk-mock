package tickets

import (
	"encoding/json"
	"strconv"

	"github.com/AirHelp/zendesk-mock/utils"
	"github.com/go-martini/martini"
)

type response struct {
	Ticket ticket `json:"ticket"`
}
type ticket struct {
	Id      int    `json:"id"`
	Subject string `json:"subject"`
}

// New creates new ticket
func New(params martini.Params) (int, string) {
	return utils.NotImplemented()
}

// Find finds ticket
func Find(params martini.Params) (int, string) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return 404, err.Error()
	}
	response := response{ticket{id, "Anything from Zendesk"}}
	bytes, err := json.Marshal(response)
	if err != nil {
		return 500, err.Error()
	}
	return 200, string(bytes)
}
