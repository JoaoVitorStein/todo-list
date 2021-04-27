package server

import (
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/todo_list/migrations"
	"github.com/todo_list/pkg/config"
)

func NewDatabase() *sqlx.DB {
	cfg := config.NewConfig()

	db, err := sqlx.Connect(cfg.DatabaseDriver, cfg.DatabaseConnectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	migrateDb(cfg.DatabaseConnectionString)
	return db
}

func migrateDb(dbUrl string) {
	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, dbUrl)
	if err != nil {
		panic(err)
	}
	m.Up() // run your migrations and handle the errors above of course
}
