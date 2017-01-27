package tickets

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"strings"
)

const (
	// TicketsFindURI uri path without param
	TicketsFindURI = "/api/v2/tickets/"
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
func Find(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(strings.Replace(req.URL.RequestURI(), TicketsFindURI, "", 1))
	if err != nil {
		log.Print(err)
		res.WriteHeader(404)
		return
	}
	response := response{ticket{id, "Anything from Zendesk", ""}}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
		res.WriteHeader(500)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(200)
	res.Write(bytes)
}
