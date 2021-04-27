package todo_list_test

import (
	"testing"

	"github.com/todo_list/pkg/server"
	"github.com/todo_list/pkg/todo_list"
)

func TestRepositoryGetById(t *testing.T) {
	db := server.NewDatabase()
	r := todo_list.NewRepository(db)

	result, err := r.GetById()

	if err != nil {
		t.Errorf("repository.getById threw an error '%v'", err)
	}

	if result != 1 {
		t.Errorf("repository.getById got = '%v', want '%v'", result, 1)
	}
}
