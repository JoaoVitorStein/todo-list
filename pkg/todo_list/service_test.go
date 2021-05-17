package todo_list_test

import (
	"testing"

	"github.com/todo_list/pkg/todo_list"
)

var getByIdMock func(int) (*todo_list.TodoListEntity, error)

type mockRepository struct {
}

func (m *mockRepository) GetById(id int) (*todo_list.TodoListEntity, error) {
	return getByIdMock(id)
}

func TestServiceGetById(t *testing.T) {
	s := todo_list.NewService(&mockRepository{})

	response := &todo_list.TodoListEntity{Id: 1, Description: "test", Done: false}
	getByIdMock = func(id int) (*todo_list.TodoListEntity, error) {
		return response, nil
	}
	result, _ := s.GetById(1)

	if result != response {
		t.Errorf("repository.getById got = '%v', want '%v'", result, response)
	}
}
