package todo_list_test

import (
	"reflect"
	"testing"

	"github.com/todo_list/pkg/server"
	"github.com/todo_list/pkg/todo_list"
)

var db, _ = server.NewDatabase()
var r = todo_list.NewRepository(db)

func TestRepositoryGetById(t *testing.T) {
	item := todo_list.TodoListEntity{
		Description: "test", Done: false,
	}

	id, _ := r.Save(item)
	item.Id = id
	result, err := r.GetById(item.Id)

	if err != nil {
		t.Errorf("Unexpected error while getting by id: %v", err)
	}

	if !reflect.DeepEqual(&item, result) {
		t.Errorf("repository.getById got = '%v', want '%v'", result, item)
	}

	r.Delete((item.Id))
}

func TestRepositorySave(t *testing.T) {
	item := todo_list.TodoListEntity{
		Description: "test", Done: false,
	}

	id, err := r.Save(item)

	if err != nil {
		t.Errorf("Unexpected error while saving: %v", err)
	}

	if id == 0 {
		t.Errorf("repository.save got = '%v', want value > 0", id)
	}

	r.Delete((item.Id))
}

func TestRepositoryDelete(t *testing.T) {
	item := todo_list.TodoListEntity{
		Description: "test", Done: false,
	}

	id, _ := r.Save(item)

	err := r.Delete(id)

	if err != nil {
		t.Errorf("Unexpected error while deleting: %v", err)
	}

	existingItem, _ := r.GetById(id)

	if existingItem.Id > 0 {
		t.Errorf("repository.delete got = '%v', want 0", existingItem.Id)
	}
}
