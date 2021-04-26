package server

import (
	"log"
	"os"

	"github.com/http-service/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase() *sqlx.DB {
	cfg := config.NewConfig()

	db, err := sqlx.Connect(cfg.DatabaseDriver, cfg.DatabaseConnectionString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	return db
}
