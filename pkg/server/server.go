package server

import (
	"github.com/gorilla/mux"
	"github.com/http-service/pkg/todo_list"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	Db       *sqlx.DB
	Router   *mux.Router
	todoList *todo_list.TodoList
}

func NewServer() *Server {
	db := NewDatabase()
	server := &Server{
		Db: db, Router: mux.NewRouter(), todoList: todo_list.NewTodoList(db),
	}
	server.routes()
	return server
}
