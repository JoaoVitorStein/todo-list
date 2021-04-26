package server

func (s *Server) routes() {
	s.Router.HandleFunc("/hello-world", s.todoList.HandleGetById())
}
