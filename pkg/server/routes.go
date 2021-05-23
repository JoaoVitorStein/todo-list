package server

func (s Server) routes() {
	s.Router.HandleFunc("/todo-list/{id:[0-9]+}", s.handleTodoListGetById()).Methods("GET")
	s.Router.HandleFunc("/todo-list", s.handleTodoListSave()).Methods("POST")
	s.Router.HandleFunc("/todo-list/{id:[0-9]+}", s.handleTodoListDelete()).Methods("DELETE")
}
