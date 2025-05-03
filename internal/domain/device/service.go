package device

type Service interface {
	CreateDevice(input *Device) error
	GetDeviceByID(id int) (*Device, error)
	GetAllDevices() ([]*Device, error)
	UpdateDevice(input *Device) error
	DeleteDevice(id int) error
	GetDevicesByEmployeeID(employeeID int) ([]*Device, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateDevice(input *Device) error {
	return s.repo.Create(input)
}

func (s *service) GetDeviceByID(id int) (*Device, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAllDevices() ([]*Device, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateDevice(input *Device) error {
	return s.repo.Update(input)
}

func (s *service) DeleteDevice(id int) error {
	return s.repo.Delete(id)
}

func (s *service) GetDevicesByEmployeeID(employeeID int) ([]*Device, error) {
	return s.repo.GetByEmployeeID(employeeID)
}
