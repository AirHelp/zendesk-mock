package mocks

import (
	"github.com/AirHelp/zendesk-mock/api"
	"time"
)

func Id() int {
	return int(time.Now().Unix())
}

func JobStatusMock() api.JobStatusEnvelope {
	resource := api.JobStatus{
		Id:       Id(),
		Url:      "url",
		Total:    1,
		Progress: 0,
		Status:   "queued"}
	return api.JobStatusEnvelope{resource}
}

func GroupMembershipsMock(userId int) api.GroupMembershipsEnvelope {
	collection := []api.GroupMembership{
		GroupMembershipMock(userId, 0),
		GroupMembershipMock(userId, 1)}
	return api.GroupMembershipsEnvelope{collection}
}

func GroupMembershipMock(userId int, idOffset int) api.GroupMembership {
	return api.GroupMembership{
		Id:      Id() + idOffset,
		Default: false,
		UserId:  userId,
		GroupId: Id() + idOffset}
}

func GroupMock(id int, name string) api.GroupEnvelope {
	return api.GroupEnvelope{api.Group{Id: id, Name: name}}
}
