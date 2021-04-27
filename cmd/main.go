package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/todo_list/pkg/server"
)

func main() {
	s, err := server.NewServer()

	if err != nil {
		log.Printf("Error while instantiating server: %s", err)
		os.Exit(1)
	}

	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
