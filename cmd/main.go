package main

import (
	"log"
	"net/http"

	"github.com/todo_list/pkg/server"
	_ "github.com/lib/pq"
)

func main() {
	s := server.NewServer()

	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
