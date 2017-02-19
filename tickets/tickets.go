package tickets

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// TicketsFindURI uri path without param
	TicketsFindURI = "/api/v2/tickets/"
)

type response struct {
	Ticket ticket `json:"ticket"`
}

type customField struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type ticket struct {
	ID           int           `json:"id"`
	Subject      string        `json:"subject"`
	Comment      string        `json:"comment"`
	CustomFields []customField `json:"custom_fields"`
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
	timestampID := int(time.Now().Unix())
	responseBody := response{ticket{ID: timestampID, Subject: input.Ticket.Subject, Comment: input.Ticket.Comment.Body}}
	bytes, err := json.Marshal(responseBody)
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
	customs := []customField{customField{23020926, "ch_web"}}
	builtResponse := response{ticket{id, "Anything from Zendesk", "", customs}}
	bytes, err := json.Marshal(builtResponse)
	if err != nil {
		log.Print(err)
		res.WriteHeader(500)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(200)
	res.Write(bytes)
}
