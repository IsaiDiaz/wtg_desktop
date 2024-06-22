package bl

import (
	"github.com/IsaiDiaz/db"
)

func CrearEmpleado(empleado *db.Empleado) Response {
	if err := db.CrearEmpleado(empleado); err != nil {
		return Response{Codigo: 400, Mensaje: "Error al crear el empleado", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Empleado creado correctamente", Data: empleado}
}

func ObtenerEmpleadoPorID(id uint) Response {
	empleado, err := db.ObtenerEmpleadoPorID(id)
	if err != nil {
		return Response{Codigo: 400, Mensaje: "Error al obtener el empleado", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Empleado obtenido correctamente", Data: empleado}
}

func AsignarTarjetaEmpleado(empleadoTarjeta *db.EmpleadoTarjeta) Response {
	if err := db.AsignarTarjetaEmpleado(empleadoTarjeta); err != nil {
		return Response{Codigo: 400, Mensaje: "Error al asignar la tarjeta al empleado", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Tarjeta asignada correctamente", Data: empleadoTarjeta}
}

func ObtenerTarjetaEmpleado(empid uint) Response {
	empleadoTarjeta, err := db.ObtenerTarjetaEmpleado(empid)
	if err != nil {
		return Response{Codigo: 400, Mensaje: "Error al obtener la tarjeta del empleado", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Tarjeta obtenida correctamente", Data: empleadoTarjeta}
}
