package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	// UsersFindPath URi path without param
	UsersFindPath = "/api/v2/users/"
)

type response struct {
	User user `json:"user"`
}

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Find finds users
func Find(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(strings.Replace(req.URL.RequestURI(), UsersFindPath, "", 1))
	if err != nil {
		log.Print(err)
		res.WriteHeader(400)
		return
	}
	user := response{user{id, "Name from Zendesk"}}
	bytes, err := json.Marshal(user)
	if err != nil {
		log.Print(err)
		res.WriteHeader(500)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(200)
	res.Write(bytes)
}
