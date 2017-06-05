package main

import (
	"fmt"
	"os"

	"github.com/AirHelp/zendesk-mock/groups"
	"github.com/AirHelp/zendesk-mock/tickets"
	"github.com/AirHelp/zendesk-mock/users"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	m.Post("/api/v2/tickets", tickets.New)
	m.Get(tickets.TicketsFindURI+":id", tickets.Find)
	m.Put(tickets.TicketsFindURI+":id", tickets.Find)

	m.Post("/api/v2/groups", groups.Create)
	m.Get(groups.ApiUrl+":id", groups.Show)
	m.Put(groups.ApiUrl+":id", groups.Update)

	m.Get(users.UsersFindPath+":id", users.Find)
	port := ":8080"
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}
	m.RunOnAddr(port)
}
