package tickets

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"time"
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
		log.Print(err)
		res.WriteHeader(400)
		return
	}
	if input.Ticket.Subject == "" {
		log.Println("Missing ticket.subject")
	}
	timestampId := int(time.Now().Unix())
	response := response{ticket{ID: timestampId, Subject: input.Ticket.Subject, Comment: input.Ticket.Comment.Body}}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
		res.WriteHeader(500)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(201)
	res.Write(bytes)

}

// Find finds ticket
func Find(params martini.Params) (int, string) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Print(err)
		return 404, err.Error()
	}
	response := response{ticket{id, "Anything from Zendesk", ""}}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
		return 500, err.Error()
	}
	return 200, string(bytes)
}
