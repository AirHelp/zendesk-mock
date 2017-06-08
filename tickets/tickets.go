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

type Response struct {
	Ticket Ticket `json:"ticket"`
}

// CustomField is used as a part of Ticket struct
type CustomField struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Ticket is used for a response
type Ticket struct {
	ID           int           `json:"id"`
	Subject      string        `json:"subject"`
	Comment      string        `json:"comment"`
	CustomFields []CustomField `json:"custom_fields"`
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

// New creates new ticket
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
	responseBody := Response{Ticket{ID: timestampID, Subject: input.Ticket.Subject, Comment: input.Ticket.Comment.Body}}
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
	customs := []CustomField{
		CustomField{28367069, "GF4534"},                                                //reference
		CustomField{28518725, "delayed"},                                               //type
		CustomField{28367089, "600"},                                                   //comp_per_pass
		CustomField{28518845, "600.0"},                                                 //comp_per_pass_decimal
		CustomField{28518825, "600"},                                                   //comp_total
		CustomField{28367689, "600.0"},                                                 //compt_total_decimal
		CustomField{28367989, "en"},                                                    //locale
		CustomField{28368009, "ch_web"},                                                //channel
		CustomField{28518885, "John Doe"},                                              //pass1
		CustomField{28518905, "John F. Kennedy International Airport, New York (JFK)"}, //departure
		CustomField{28368029, "Tegel Airport, Berlin (TXL)"},                           //arrival
		CustomField{28518945, "LOT - Polish Airlines (LO)"},                            //airline_name_and_code
		CustomField{28368069, "554"},                                                   //flight_no
		CustomField{28368089, "2017-03-13"},                                            //date
		CustomField{28518965, "LO"},                                                    //airline_code
		CustomField{29960329, ""}}                                                      //linked_ticket

	builtResponse := Response{Ticket{id, "Anything from Zendesk", "", customs}}
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
