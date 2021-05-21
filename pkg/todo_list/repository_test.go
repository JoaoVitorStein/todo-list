package todo_list_test

import (
	"reflect"
	"testing"

	"github.com/todo_list/pkg/server"
	"github.com/todo_list/pkg/todo_list"
)

func TestRepositoryGetById(t *testing.T) {
	db, _ := server.NewDatabase()
	r := todo_list.NewRepository(db)
	item := &todo_list.TodoListEntity{
		Id: 1, Description: "test", Done: false,
	}

	db.NamedExec("INSERT INTO todo_list(id, description, done) VALUES (:id, :description, :done)", &item)
	result, _ := r.GetById(item.Id)

	if !reflect.DeepEqual(item, result) {
		t.Errorf("repository.getById got = '%v', want '%v'", result, item)
	}
}
