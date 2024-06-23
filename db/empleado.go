package db

func CrearEmpleado(empleado *Empleado) error {
	return DB.Create(empleado).Error
}

func ObtenerEmpleadoPorID(id uint) (*Empleado, error) {
	var empleado Empleado
	err := DB.Model(&Empleado{}).Preload("Tarjeta").Where("empleado_id = ?", id).First(&empleado).Error
	if err != nil {
		return nil, err
	}
	return &empleado, nil
}

func ObtenerTodosEmpleados() ([]Empleado, error) {
	var empleados []Empleado
	err := DB.Model(&Empleado{}).Preload("Tarjeta").Find(&empleados).Error
	if err != nil {
		return nil, err
	}
	return empleados, nil
}

func AsignarTarjetaEmpleado(empid uint, tarid uint) error {
	var empleado Empleado
	var tarjeta Tarjeta
	err := DB.First(&empleado, empid).Error
	if err != nil {
		return err
	}
	err = DB.First(&tarjeta, tarid).Error
	if err != nil {
		return err
	}

	empleado.Tarjeta = tarjeta
	return DB.Save(&empleado).Error
}

func ObtenerTarjetaEmpleado(empid uint) (*Empleado, error) {
	var empleado Empleado
	err := DB.Model(&Empleado{}).Preload("Tarjeta").Where("empleado_id = ?", empid).First(&empleado).Error
	if err != nil {
		return nil, err
	}
	return &empleado, nil
}
