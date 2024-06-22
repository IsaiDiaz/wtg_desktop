package db

func CrearEmpleado(empleado *Empleado) error {
	return DB.Create(empleado).Error
}

func ObtenerEmpleadoPorID(id uint) (*Empleado, error) {
	var empleado Empleado
	err := DB.First(&empleado, id).Error
	if err != nil {
		return nil, err
	}
	return &empleado, nil
}

func AsignarTarjetaEmpleado(empleadoTarjeta *EmpleadoTarjeta) error {
	return DB.Create(empleadoTarjeta).Error
}

func ObtenerTarjetaEmpleado(empid uint) (*Tarjeta, error) {
	var empleadoTarjeta EmpleadoTarjeta
	err := DB.Where("empleado_id = ?", empid).First(&empleadoTarjeta).Error
	if err != nil {
		return nil, err
	}
	var tarjeta Tarjeta
	err = DB.First(&tarjeta, empleadoTarjeta.TarjetaID).Error
	if err != nil {
		return nil, err
	}
	return &tarjeta, nil
}
