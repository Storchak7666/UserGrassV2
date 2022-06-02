package desease_list

type Service interface {
	FindAll() ([]DeseaseList, error)
	FindById(id int64) (*DeseaseList, error)
	CreateDesease(user DeseaseList) (int64, error)
	//UpdateById(user DeseaseList) (*DeseaseList, error)
}

type service struct {
	repo *DeseaseListRepository
}

func NewDeseaseListService(r *DeseaseListRepository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]DeseaseList, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindById(id int64) (*DeseaseList, error) {
	return (*s.repo).FindById(id)
}

func (s *service) CreateDesease(user DeseaseList) (int64, error) {
	return (*s.repo).CreateDesease(user)
}

//func (s *service) UpdateById(user DeseaseList) (*DeseaseList, error) {
//	return (*s.repo).UpdateById(user)
//}
