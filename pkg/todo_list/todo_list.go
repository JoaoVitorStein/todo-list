package todo_list

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type TodoList struct {
	service *Service
}

func NewTodoList(db *sqlx.DB) *TodoList {
	return &TodoList{
		service: NewService(NewRepository(db)),
	}
}

func (t *TodoList) HandleGetById() http.HandlerFunc {
	type response struct {
		Id int `json:"id"`
	}
	type errorResponse struct {
		Message string `json:"message"`
	}
	return func(res http.ResponseWriter, req *http.Request) {
		result, err := t.service.GetById()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(&errorResponse{
				Message: "Internal server error",
			})
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(&response{
			Id: result,
		})
	}
}
