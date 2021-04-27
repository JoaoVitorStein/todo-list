package server

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/todo_list/pkg/todo_list"
)

type Server struct {
	db       *sqlx.DB
	Router   *mux.Router
	todoList *todo_list.TodoList
}

func NewServer() *Server {
	db := NewDatabase()
	server := &Server{
		db: db, Router: mux.NewRouter(), todoList: todo_list.NewTodoList(db),
	}
	server.routes()
	return server
}
