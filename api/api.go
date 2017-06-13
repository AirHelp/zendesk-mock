package api

type JobStatusEnvelope struct {
	JobStatus JobStatus `json:"job_status"`
}

type JobStatus struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	Total    int    `json:"total"`
	Progress int    `json:"progress"`
	Status   string `json:"status"`
}

type GroupMembership struct {
	Id      int  `json:"id"`
	Default bool `json:"default"`
	UserId  int  `json:"user_id"`
	GroupId int  `json:"group_id"`
}

type GroupMembershipsEnvelope struct {
	GroupMemberships []GroupMembership `json:"group_memberships"`
}

type GroupEnvelope struct {
	Group Group `json:"group"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserEnvelope struct {
	User User `json:"user"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
