package mocks

import (
	"time"

	"github.com/AirHelp/zendesk-mock/api"
)

func Id() int {
	return int(time.Now().Unix())
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
		api.CustomField{28367069, "GF4534"},                                                //reference
		api.CustomField{28518725, "delayed"},                                               //type
		api.CustomField{28367089, "600"},                                                   //comp_per_pass
		api.CustomField{28518845, "600.0"},                                                 //comp_per_pass_decimal
		api.CustomField{28518825, "600"},                                                   //comp_total
		api.CustomField{28367689, "600.0"},                                                 //compt_total_decimal
		api.CustomField{28367989, "en"},                                                    //locale
		api.CustomField{28368009, "ch_web"},                                                //channel
		api.CustomField{28518885, "John Doe"},                                              //pass1
		api.CustomField{28518905, "John F. Kennedy International Airport, New York (JFK)"}, //departure
		api.CustomField{28368029, "Tegel Airport, Berlin (TXL)"},                           //arrival
		api.CustomField{28518945, "LOT - Polish Airlines (LO)"},                            //airline_name_and_code
		api.CustomField{28368069, "554"},                                                   //flight_no
		api.CustomField{28368089, "2017-03-13"},                                            //date
		api.CustomField{28518965, "LO"},                                                    //airline_code
		api.CustomField{29960329, ""},                                                      //linked_ticket
	}
	return api.TicketEnvelope{api.Ticket{id, subject, customs}}
}

func User(id int, name string) api.UserEnvelope {
	return api.UserEnvelope{User: api.User{Id: id, Name: name}}
}
