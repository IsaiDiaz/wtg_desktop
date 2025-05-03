package rfidcard

type Service interface {
	CreateCard(card *RfidCard) error
	GetAllCards() ([]*RfidCard, error)
	GetCardByUUID(uuid string) (*RfidCard, error)
	GetCardByID(id uint) (*RfidCard, error)
	UpdateCard(card *RfidCard) error
	DeleteCard(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateCard(card *RfidCard) error {
	return s.repo.Create(card)
}

func (s *service) GetAllCards() ([]*RfidCard, error) {
	return s.repo.GetAll()
}

func (s *service) GetCardByUUID(uuid string) (*RfidCard, error) {
	return s.repo.GetByUUID(uuid)
}

func (s *service) GetCardByID(id uint) (*RfidCard, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateCard(card *RfidCard) error {
	return s.repo.Update(card)
}

func (s *service) DeleteCard(id uint) error {
	return s.repo.Delete(id)
}
