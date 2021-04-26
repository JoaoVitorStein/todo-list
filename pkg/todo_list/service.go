package todo_list

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetById() (int, error) {
	return s.repository.GetById()
}
