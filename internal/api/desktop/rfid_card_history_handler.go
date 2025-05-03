package desktop

import "wtg_desktop/internal/domain/rfidcardhistory"

type RfidCardHistoryHandler struct {
	service rfidcardhistory.Service
}

func NewRfidCardHistoryHandler(service rfidcardhistory.Service) *RfidCardHistoryHandler {
	return &RfidCardHistoryHandler{service: service}
}

func (h *RfidCardHistoryHandler) RegisterAssignment(record *rfidcardhistory.RfidCardHistory) error {
	return h.service.RegisterAssignment(record)
}

func (h *RfidCardHistoryHandler) GetHistoryForEmployee(employeeID uint) ([]rfidcardhistory.RfidCardHistory, error) {
	return h.service.GetHistoryForEmployee(employeeID)
}
