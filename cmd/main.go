package main

import (
	"log"
	"net/http"

	"github.com/http-service/pkg/server"
	_ "github.com/lib/pq"
)

func main() {
	s := server.NewServer()

	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
