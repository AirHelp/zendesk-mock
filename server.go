package main

import (
	"fmt"
	"os"

	"github.com/AirHelp/zendesk-mock/group_memberships"
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
	m.Get("/api/v2/groups/:id", groups.Show)
	m.Put("/api/v2/groups/:id", groups.Update)

	m.Get("/api/v2/users/:user_id/group_memberships", group_memberships.Index)
	m.Delete("/api/v2/group_memberships/destroy_many.json", group_memberships.DestroyMany)
	m.Post("/api/v2/group_memberships/create_many.json", group_memberships.CreateMany)

	m.Get(users.UsersFindPath+":id", users.Find)

	port := ":8080"
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}
	m.RunOnAddr(port)
}
