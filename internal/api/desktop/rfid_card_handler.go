package desktop

import "wtg_desktop/internal/domain/rfidcard"

type RfidCardHandler struct {
	service rfidcard.Service
}

func NewRfidCardHandler(service rfidcard.Service) *RfidCardHandler {
	return &RfidCardHandler{service: service}
}

func (r *RfidCardHandler) CreateRfidCard(card *rfidcard.RfidCard) error {
	return r.service.CreateCard(card)
}

func (r *RfidCardHandler) GetRfidCardByID(id uint) (*rfidcard.RfidCard, error) {
	return r.service.GetCardByID(id)
}

func (r *RfidCardHandler) GetAllRfidCards() ([]*rfidcard.RfidCard, error) {
	return r.service.GetAllCards()
}

func (r *RfidCardHandler) UpdateRfidCard(card *rfidcard.RfidCard) error {
	return r.service.UpdateCard(card)
}

func (r *RfidCardHandler) DeleteRfidCard(id uint) error {
	return r.service.DeleteCard(id)
}

func (r *RfidCardHandler) GetRfidCardByUUID(uuid string) (*rfidcard.RfidCard, error) {
	return r.service.GetCardByUUID(uuid)
}
