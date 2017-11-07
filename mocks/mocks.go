package mocks

import (
	"math/rand"
	"time"

	"github.com/AirHelp/zendesk-mock/api"
)

func Id() int {
	return int(time.Now().Unix()) + rand.Intn(10000)
}

func JobStatus() api.JobStatusEnvelope {
	resource := api.JobStatus{
		Id:       Id(),
		Url:      "url",
		Total:    1,
		Progress: 0,
		Status:   "queued"}
	return api.JobStatusEnvelope{JobStatus: resource}
}

func GroupMemberships(userId int) api.GroupMembershipsEnvelope {
	collection := []api.GroupMembership{
		GroupMembership(userId, 0),
		GroupMembership(userId, 1)}
	return api.GroupMembershipsEnvelope{GroupMemberships: collection}
}

func GroupMembership(userId int, idOffset int) api.GroupMembership {
	return api.GroupMembership{
		Id:      Id() + idOffset,
		Default: false,
		UserId:  userId,
		GroupId: Id() + idOffset}
}

func Group(id int, name string) api.GroupEnvelope {
	return api.GroupEnvelope{Group: api.Group{Id: id, Name: name}}
}

func Ticket(id int, subject string) api.TicketEnvelope {
	customs := []api.CustomField{
		api.CustomField{
			ID:    28367069,
			Value: "GF4534",
		}, //reference
		api.CustomField{
			ID:    28518725,
			Value: "delayed",
		}, //type
		api.CustomField{
			ID:    28367089,
			Value: "600",
		}, //comp_per_pass
		api.CustomField{
			ID:    28518845,
			Value: "600.0",
		}, //comp_per_pass_decimal
		api.CustomField{
			ID:    28518825,
			Value: "600",
		}, //comp_total
		api.CustomField{
			ID:    28367689,
			Value: "600.0",
		}, //compt_total_decimal
		api.CustomField{
			ID:    28367989,
			Value: "en",
		}, //locale
		api.CustomField{
			ID:    28368009,
			Value: "ch_web",
		}, //channel
		api.CustomField{
			ID:    28518885,
			Value: "John Doe",
		}, //pass1
		api.CustomField{
			ID:    28518905,
			Value: "John F. Kennedy International Airport, New York (JFK)",
		}, //departure
		api.CustomField{
			ID:    28368029,
			Value: "Tegel Airport, Berlin (TXL)",
		}, //arrival
		api.CustomField{
			ID:    28518945,
			Value: "LOT - Polish Airlines (LO)",
		}, //airline_name_and_code
		api.CustomField{
			ID:    28368069,
			Value: "554",
		}, //flight_no
		api.CustomField{
			ID:    28368089,
			Value: "2017-03-13",
		}, //date
		api.CustomField{
			ID:    28518965,
			Value: "LO",
		}, //airline_code
		api.CustomField{
			ID:    29960329,
			Value: "",
		}, //linked_ticket
	}
	return api.TicketEnvelope{
		Ticket: api.Ticket{
			ID:           id,
			Subject:      subject,
			CustomFields: customs,
		},
	}
}

func User(id int, name string) api.UserEnvelope {
	return api.UserEnvelope{User: api.User{Id: id, Name: name}}
}
