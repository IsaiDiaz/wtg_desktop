package category

type Service interface {
	CreateCategory(input *Category) error
	GetCategoryByID(id uint) (*Category, error)
	GetAllCategories() ([]*Category, error)
	UpdateCategory(input *Category) error
	DeleteCategory(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateCategory(input *Category) error {
	return s.repo.Create(input)
}

func (s *service) GetCategoryByID(id uint) (*Category, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAllCategories() ([]*Category, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateCategory(input *Category) error {
	return s.repo.Update(input)
}

func (s *service) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}
