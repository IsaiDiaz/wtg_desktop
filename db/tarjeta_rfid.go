package db

func CrearTarjetaRFID(tarjeta *Tarjeta) error {
	return DB.Create(tarjeta).Error
}

func ObtenerTarjetaRFIDPorID(id uint) (*Tarjeta, error) {
	var tarjeta Tarjeta
	err := DB.First(&tarjeta, id).Error
	if err != nil {
		return nil, err
	}
	return &tarjeta, nil
}

func ObtenerTarjetaRFIDPorUID(uid string) (*Tarjeta, error) {
	var tarjeta Tarjeta
	err := DB.Where("codigo_tarjeta = ?", uid).First(&tarjeta).Error
	if err != nil {
		return nil, err
	}
	return &tarjeta, nil
}

func ObtenerTodasTarjetasRFID() ([]Tarjeta, error) {
	var tarjetas []Tarjeta
	err := DB.Find(&tarjetas).Error
	if err != nil {
		return nil, err
	}
	return tarjetas, nil
}
