package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (t *Server) handleTodoListGetById() http.HandlerFunc {
	type response struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
		Done        bool   `json:"done"`
	}
	type errorResponse struct {
		Message string `json:"message"`
	}
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(&errorResponse{
				Message: "Internal server error",
			})
		}
		result, err := t.todoList.GetById(id)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(&errorResponse{
				Message: "Internal server error",
			})
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(&response{
			Id:          result.Id,
			Description: result.Description,
			Done:        result.Done,
		})
	}
}
