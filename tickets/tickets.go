package tickets

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

const (
	min = 1000
	max = 10000
)

type response struct {
	Ticket ticket `json:"ticket"`
}

type ticket struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	Comment string `json:"comment"`
}

type reqBody struct {
	Ticket reqTicket `json:"ticket"`
}
type reqTicket struct {
	Subject string     `json:"subject"`
	Comment reqComment `json:"comment"`
}
type reqComment struct {
	Body string `json:"body"`
}

func New(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var input reqBody
	if err := dec.Decode(&input); err != nil {
		res.WriteHeader(400)
		return
	}
	if input.Ticket.Subject == "" {
		res.WriteHeader(400)
		return
	}
	response := response{ticket{ID: min + rand.Intn(max), Subject: input.Ticket.Subject, Comment: input.Ticket.Comment.Body}}
	bytes, err := json.Marshal(response)
	if err != nil {
		res.WriteHeader(500)
		return
	}
	res.WriteHeader(201)
	res.Write(bytes)

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
