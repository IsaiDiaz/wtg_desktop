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
