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

type TicketEnvelope struct {
	Ticket Ticket `json:"ticket"`
}

type Ticket struct {
	ID           int           `json:"id"`
	Subject      string        `json:"subject"`
	CustomFields []CustomField `json:"custom_fields"`
}

type CustomField struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type CreateTicketEnvelope struct {
	Ticket CreateTicket `json:"ticket"`
}

type CreateTicket struct {
	Subject string  `json:"subject"`
	Comment Comment `json:"comment"`
}

type Comment struct {
	Body string `json:"body"`
}
