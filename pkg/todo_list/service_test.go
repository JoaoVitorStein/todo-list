package todo_list_test

import (
	"testing"

	"github.com/todo_list/pkg/todo_list"
)

var getByIdMock func(int) (*todo_list.TodoListEntity, error)
var saveMock func(todo_list.TodoListEntity) (int, error)
var deleteMock func(int) error

type mockRepository struct {
}

func (m mockRepository) GetById(id int) (*todo_list.TodoListEntity, error) {
	return getByIdMock(id)
}

func (m mockRepository) Save(data todo_list.TodoListEntity) (int, error) {
	return saveMock(data)
}

func (m mockRepository) Delete(id int) error {
	return deleteMock(id)
}

var s = todo_list.NewService(mockRepository{})

func TestServiceGetById(t *testing.T) {
	response := &todo_list.TodoListEntity{Id: 1, Description: "test", Done: false}
	getByIdMock = func(id int) (*todo_list.TodoListEntity, error) {
		return response, nil
	}
	result, _ := s.GetById(1)

	if result != response {
		t.Errorf("service.GetById got = '%v', want '%v'", result, response)
	}
}

func TestServiceSave(t *testing.T) {
	data := todo_list.TodoListEntity{Id: 1, Description: "test", Done: false}
	saveMock = func(data todo_list.TodoListEntity) (int, error) {
		return data.Id, nil
	}
	result, _ := s.Save(data)

	if result != data.Id {
		t.Errorf("service.Save got = '%v', want '%v'", result, data.Id)
	}
}

func TestServiceDelete(t *testing.T) {
	deleteMock = func(id int) error {
		return nil
	}
	err := s.Delete(1)

	if err != nil {
		t.Errorf("service.Delete got = '%v', want '%v'", err, nil)
	}
}
