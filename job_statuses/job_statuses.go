package job_statuses

import (
)

const (
	ApiUrl = "/api/v2/job_statuses/"
)

type ResourceEnvelope struct {
	 Resource Resource `json:"job_status"`
}

type Resource struct {
	Id   int  		`json:"id"`
	Url string 		`json:"url"`
	Total int 		`json:"total"`
	Progress int 	`json:"progress"`
  Status string `json:"status"`
}
