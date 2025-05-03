package desktop

import (
	"context"
	"wtg_desktop/internal/domain/employee"
)

type EmployeeHandler struct {
	service employee.Service
	ctx     context.Context
}

func NewEmployeeHandler(service employee.Service) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (e *EmployeeHandler) WailsInit(ctx context.Context) {
	e.ctx = ctx
}

func (e *EmployeeHandler) CreateEmployee(emp *employee.Employee) error {
	return e.service.CreateEmployee(emp)
}

func (e *EmployeeHandler) GetEmployeeByID(id int) (*employee.Employee, error) {
	return e.service.GetEmployeeByID(id)
}

func (e *EmployeeHandler) GetAllEmployees() ([]*employee.Employee, error) {
	return e.service.GetAllEmployees()
}

func (e *EmployeeHandler) UpdateEmployee(emp *employee.Employee) error {
	return e.service.UpdateEmployee(emp)
}

func (e *EmployeeHandler) DeleteEmployee(id int) error {
	return e.service.DeleteEmployee(id)
}
