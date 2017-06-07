package job_statuses

import (
	"encoding/json"
	"github.com/AirHelp/zendesk-mock/respond"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ApiUrl = "/api/v2/job_statuses/"
)

type resourceEnvelope struct {
	 JobStatus JobStatus `json:"job_status"`
}

type JobStatus struct {
	Id   int  		`json:"id"`
	Url string 		`json:"url"`
	Total int 		`json:"total"`
	Progress int 	`json:"progress"`
  Status string `json:"status"`
}
