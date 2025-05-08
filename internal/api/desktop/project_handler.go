package desktop

import "wtg_desktop/internal/domain/project"

type ProjectHandler struct {
	service project.Service
}

func NewProjectHandler(service project.Service) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) CreateProject(input *project.Project) error {
	return h.service.CreateProject(input)
}

func (h *ProjectHandler) GetProjectByID(id uint) (*project.Project, error) {
	return h.service.GetProjectByID(id)
}

func (h *ProjectHandler) GetAllProjects() ([]*project.Project, error) {
	return h.service.GetAllProjects()
}

func (h *ProjectHandler) UpdateProject(input *project.Project) error {
	return h.service.UpdateProject(input)
}

func (h *ProjectHandler) DeleteProject(id uint) error {
	return h.service.DeleteProject(id)
}
