package todo_list

import (
	"github.com/jmoiron/sqlx"
)

type repository interface {
	GetById(int) (*TodoListEntity, error)
	Save(TodoListEntity) (int, error)
	Delete(int) error
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

func (r repositoryPg) Save(data TodoListEntity) (int, error) {
	var id int
	row, err := r.db.NamedQuery("INSERT INTO todo_list(description, done) VALUES (:description, :done) RETURNING id", data)
	if err != nil {
		return 0, err
	}
	if row.Next() {
		row.Scan(&id)
	}
	return id, nil

}

func (r repositoryPg) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todo_list WHERE id = $1", id)
	return err
}
