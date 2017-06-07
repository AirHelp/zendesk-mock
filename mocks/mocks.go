package mocks

import (
	"time"
	"github.com/AirHelp/zendesk-mock/job_statuses"
)

type JobStatusEnvelope struct {
	 JobStatus JobStatus `json:"job_status"`
}

type JobStatus struct {
	Id   int  		`json:"id"`
	Url string 		`json:"url"`
	Total int 		`json:"total"`
	Progress int 	`json:"progress"`
  Status string `json:"status"`
}


type GroupMembership struct {
	Id   int    	`json:"id"`
	Default bool 	`json:"default"`
	UserId int    `json:"user_id"`
	GroupId int   `json:"group_id"`
}

type GroupMembershipsEnvelope struct {
	GroupMemberships []GroupMembership `json:"group_memberships"`
}

func Id() int {
	return int(time.Now().Unix())
}

func JobStatusMock() job_statuses.ResourceEnvelope {
	resource := job_statuses.Resource{
		Id: Id(),
		Url: "url",
		Total: 1,
		Progress: 0,
		Status: "queued" }
	return job_statuses.ResourceEnvelope{resource}
}

func GroupMembershipsMock(userId int) GroupMembershipsEnvelope {
	collection := []GroupMembership{
		GroupMembershipMock(userId, 0),
		GroupMembershipMock(userId, 1)}
	return GroupMembershipsEnvelope{collection}
}

func GroupMembershipMock(userId int, idOffset int) GroupMembership {
	return GroupMembership{
		Id: Id() + idOffset,
		Default: false,
		UserId: userId,
		GroupId:
		Id() + idOffset}
}
