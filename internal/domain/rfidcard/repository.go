package rfidcard

import "gorm.io/gorm"

type Repository interface {
	Create(card *RfidCard) error
	GetAll() ([]*RfidCard, error)
	GetByUUID(uuid string) (*RfidCard, error)
	GetByID(id uint) (*RfidCard, error)
	Update(card *RfidCard) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(card *RfidCard) error {
	return r.db.Create(card).Error
}

func (r *repository) GetAll() ([]*RfidCard, error) {
	var cards []*RfidCard
	err := r.db.Find(&cards).Error
	return cards, err
}

func (r *repository) GetByUUID(uuid string) (*RfidCard, error) {
	var card RfidCard
	err := r.db.Where("rfid_card_uuid = ?", uuid).First(&card).Error
	return &card, err
}

func (r *repository) GetByID(id uint) (*RfidCard, error) {
	var card RfidCard
	err := r.db.First(&card, id).Error
	return &card, err
}

func (r *repository) Update(card *RfidCard) error {
	return r.db.Save(card).Error
}

func (r *repository) Delete(id uint) error {
	var card RfidCard
	err := r.db.First(&card, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&card).Error
}
