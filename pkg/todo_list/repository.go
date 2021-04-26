package todo_list

import "github.com/jmoiron/sqlx"

type repository interface {
	GetById() (int, error)
}

type repositoryPg struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return &repositoryPg{db: db}
}

func (r *repositoryPg) GetById() (int, error) {
	row := r.db.QueryRow("SELECT 1 as test")
	var test int
	err := row.Scan(&test)
	return test, err
}
