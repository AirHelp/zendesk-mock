package mocks

import (
	"time"
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

func JobStatusMock() JobStatusEnvelope {
	resource := JobStatus{
		Id: Id(),
		Url: "url",
		Total: 1,
		Progress: 0,
		Status: "queued" }
	return JobStatusEnvelope{resource}
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
