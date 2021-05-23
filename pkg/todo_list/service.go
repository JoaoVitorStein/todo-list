package todo_list

type Service struct {
	repository repository
}

func NewService(repository repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) GetById(id int) (*TodoListEntity, error) {
	return s.repository.GetById(id)
}

func (s Service) Save(data TodoListEntity) (int, error) {
	return s.repository.Save(data)
}

func (s Service) Delete(id int) error {
	return s.repository.Delete(id)
}
