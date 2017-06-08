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
