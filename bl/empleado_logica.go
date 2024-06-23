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

func ObtenerTodosEmpleados() Response {
	empleados, err := db.ObtenerTodosEmpleados()
	if err != nil {
		return Response{Codigo: 400, Mensaje: "Error al obtener los empleados", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Empleados obtenidos correctamente", Data: empleados}
}

func AsignarTarjetaEmpleado(empid uint, tarid uint) Response {
	if err := db.AsignarTarjetaEmpleado(empid, tarid); err != nil {
		return Response{Codigo: 400, Mensaje: "Error al asignar la tarjeta al empleado", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Tarjeta asignada correctamente"}
}

func ObtenerTarjetaEmpleado(empid uint) Response {
	empleado, err := db.ObtenerTarjetaEmpleado(empid)
	if err != nil {
		return Response{Codigo: 400, Mensaje: "Error al obtener la tarjeta del empleado", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Tarjeta obtenida correctamente", Data: empleado}
}
