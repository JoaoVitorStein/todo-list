package todo_list_test

import (
	"testing"

	"github.com/http-service/pkg/todo_list"
)

var getByIdMock func() (int, error)

type mockRepository struct {
}

func (m *mockRepository) GetById() (int, error) {
	return getByIdMock()
}

func TestServiceGetById(t *testing.T) {
	s := todo_list.NewService(&mockRepository{})

	getByIdMock = func() (int, error) {
		return 1, nil
	}
	result, _ := s.GetById()

	if result != 1 {
		t.Errorf("repository.getById got = '%v', want '%v'", result, 1)
	}
}
