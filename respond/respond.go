package respond

import (
	"log"
	"net/http"
)

func WithJson(res http.ResponseWriter, code int, body []byte, err error) {
	if err != nil {
		log.Print(err)
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(code)
	if body != nil {
		res.Write(body)
	}
}
