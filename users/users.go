package users

import (
	"encoding/json"
	"strconv"

	"github.com/go-martini/martini"
)

type response struct {
	User user `json:"user"`
}

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Find finds users
func Find(params martini.Params) (int, string) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return 400, err.Error()
	}
	response := response{user{id, "Name from Zendesk"}}
	bytes, err := json.Marshal(response)
	if err != nil {
		return 500, err.Error()
	}
	return 200, string(bytes)
}
