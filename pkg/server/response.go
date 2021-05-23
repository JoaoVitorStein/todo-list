package server

import (
	"encoding/json"
	"net/http"
)

func (s Server) makeResponse(res http.ResponseWriter, statusCode int, body interface{}) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(body)
}
