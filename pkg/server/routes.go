package server

func (s Server) routes() {
	s.Router.HandleFunc("/todo-list/{id:[0-9]+}", s.handleTodoListGetById()).Methods("GET")
}
