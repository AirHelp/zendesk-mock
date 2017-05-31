package groups

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	GroupsFindURI = "/api/v2/groups/"
)

type Response struct {
	Group Group `json:"group"`
}

// Group is used for a response
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type reqBody struct {
	Group reqGroup `json:"group"`
}

type reqGroup struct {
	Name string `json:"name"`
}

func Create(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var input reqBody
	if err := dec.Decode(&input); err != nil {
		log.Print(err)
		res.WriteHeader(400)
		return
	}
	timestampID := int(time.Now().Unix())
	responseBody := Response{Group{ID: timestampID, Name: input.Group.Name}}
	bytes, err := json.Marshal(responseBody)
	if err != nil {
		log.Print(err)
		res.WriteHeader(500)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(201)
	res.Write(bytes)
}
