package desktop

import "wtg_desktop/internal/domain/projectemployee"

type ProjectEmployeeHandler struct {
	service projectemployee.Service
}

func NewProjectEmployeeHandler(service projectemployee.Service) *ProjectEmployeeHandler {
	return &ProjectEmployeeHandler{service: service}
}

func (p *ProjectEmployeeHandler) CreateProjectEmployee(pe *projectemployee.ProjectEmployee) error {
	return p.service.CreateProjectEmployee(pe)
}

func (p *ProjectEmployeeHandler) GetProjectEmployeeByID(id int) (*projectemployee.ProjectEmployee, error) {
	return p.service.GetProjectEmployeeByID(id)
}

func (p *ProjectEmployeeHandler) GetAllProjectEmployees() ([]*projectemployee.ProjectEmployee, error) {
	return p.service.GetAllProjectEmployees()
}

func (p *ProjectEmployeeHandler) UpdateProjectEmployee(pe *projectemployee.ProjectEmployee) error {
	return p.service.UpdateProjectEmployee(pe)
}

func (p *ProjectEmployeeHandler) DeleteProjectEmployee(id int) error {
	return p.service.DeleteProjectEmployee(id)
}

func (p *ProjectEmployeeHandler) GetProjectEmployeesByProjectID(projectID int) ([]*projectemployee.ProjectEmployee, error) {
	return p.service.GetProjectEmployeesByProjectID(projectID)
}

func (p *ProjectEmployeeHandler) GetProjectEmployeesByEmployeeID(employeeID int) ([]*projectemployee.ProjectEmployee, error) {
	return p.service.GetProjectEmployeesByEmployeeID(employeeID)
}
