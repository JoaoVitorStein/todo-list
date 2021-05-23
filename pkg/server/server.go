package server

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/todo_list/pkg/todo_list"
)

type Server struct {
	db       *sqlx.DB
	Router   *mux.Router
	todoList todo_list.Service
}

func NewServer() (*Server, error) {
	db, err := NewDatabase()
	if err != nil {
		return nil, err
	}
	server := &Server{
		db: db, Router: mux.NewRouter(), todoList: todo_list.NewService(todo_list.NewRepository(db)),
	}
	server.routes()
	return server, nil
}
