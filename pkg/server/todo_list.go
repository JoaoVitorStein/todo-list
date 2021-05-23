package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/todo_list/pkg/todo_list"
)

func (s Server) handleTodoListGetById() http.HandlerFunc {
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
			response := &errorResponse{
				Message: err.Error(),
			}
			s.makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		result, err := s.todoList.GetById(id)
		if err != nil {
			response := &errorResponse{
				Message: err.Error(),
			}
			s.makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		response := &response{
			Id:          result.Id,
			Description: result.Description,
			Done:        result.Done,
		}
		s.makeResponse(res, http.StatusOK, response)
	}
}

func (s Server) handleTodoListSave() http.HandlerFunc {
	type request struct {
		Description string `json:"description"`
		Done        bool   `json:"done"`
	}
	type response struct {
		Id int `json:"id"`
	}
	type errorResponse struct {
		Message string `json:"message"`
	}
	return func(res http.ResponseWriter, req *http.Request) {
		var body request
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			response := &errorResponse{
				Message: err.Error(),
			}
			s.makeResponse(res, http.StatusBadRequest, response)
			return
		}
		data := todo_list.TodoListEntity{
			Description: body.Description,
			Done:        body.Done,
		}
		result, err := s.todoList.Save(data)
		if err != nil {
			response := &errorResponse{
				Message: err.Error(),
			}
			s.makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		response := &response{
			Id: result,
		}
		s.makeResponse(res, http.StatusOK, response)
	}
}

func (s Server) handleTodoListDelete() http.HandlerFunc {
	type response struct {
		Ok bool `json:"ok"`
	}
	type errorResponse struct {
		Message string `json:"message"`
	}
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response := &errorResponse{
				Message: err.Error(),
			}
			s.makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		err = s.todoList.Delete(id)
		if err != nil {
			response := &errorResponse{
				Message: err.Error(),
			}
			s.makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		response := &response{
			Ok: true,
		}
		s.makeResponse(res, http.StatusOK, response)
	}
}
