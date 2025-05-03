package rfidcardhistory

type Service interface {
	RegisterAssignment(record *RfidCardHistory) error
	GetHistoryForEmployee(employeeID uint) ([]RfidCardHistory, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) RegisterAssignment(record *RfidCardHistory) error {
	return s.repo.Create(record)
}

func (s *service) GetHistoryForEmployee(employeeID uint) ([]RfidCardHistory, error) {
	return s.repo.GetByEmployeeID(employeeID)
}
