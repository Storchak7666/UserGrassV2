package user

type Service interface {
	FindAll() ([]User, error)
	FindById(id int64) (*User, error)
	CreateUser(user User) (int64, error)
	UpdateById(user User) (*User, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]User, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindById(id int64) (*User, error) {
	return (*s.repo).FindById(id)
}

func (s *service) CreateUser(user User) (int64, error) {
	return (*s.repo).CreateUser(user)
}

func (s *service) UpdateById(user User) (*User, error) {
	return (*s.repo).UpdateById(user)
}
