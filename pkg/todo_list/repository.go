package todo_list

import "github.com/jmoiron/sqlx"

type repository interface {
	GetById(int) (*TodoListEntity, error)
}

type repositoryPg struct {
	db *sqlx.DB
}

var _ repository = NewRepository(nil)

type TodoListEntity struct {
	Id          int
	Description string
	Done        bool
}

func NewRepository(db *sqlx.DB) repository {
	return repositoryPg{db: db}
}

func (r repositoryPg) GetById(id int) (*TodoListEntity, error) {
	var todoListItem TodoListEntity
	err := r.db.Get(&todoListItem, "SELECT * FROM todo_list WHERE id = $1", id)
	return &todoListItem, err
}
