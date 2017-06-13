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

	m.Group("/api/v2/tickets", func(r martini.Router) {
		r.Post("", tickets.Create)
		r.Get("/:id", tickets.Show)
		r.Put("/:id", tickets.Update)
	})

	m.Group("/api/v2/groups", func(r martini.Router) {
		r.Post("", groups.Create)
		r.Get("/:id", groups.Show)
		r.Put("/:id", groups.Update)
	})

	m.Group("/api/v2/group_memberships", func(r martini.Router) {
		r.Delete("/destroy_many", group_memberships.DestroyMany)
		r.Post("/create_many", group_memberships.CreateMany)
	})

	m.Group("/api/v2/users", func(r martini.Router) {
		r.Get("/:id", users.Show)
		r.Get("/:user_id/group_memberships", group_memberships.Index)
	})

	port := ":8080"
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}
	m.RunOnAddr(port)
}
