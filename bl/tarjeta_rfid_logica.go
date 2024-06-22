package bl

import (
	"github.com/IsaiDiaz/db"
)

func CrearTarjetaRFID(tarjeta *db.Tarjeta) Response {
	existeTarjeta, err := db.ObtenerTarjetaRFIDPorUID(tarjeta.CodigoTarjeta)

	if err != nil && existeTarjeta != nil {
		return Response{Codigo: 400, Mensaje: "Ya existe una tarjeta con ese código", Data: nil}
	}

	if err := db.CrearTarjetaRFID(tarjeta); err != nil {
		return Response{Codigo: 400, Mensaje: "Error al crear la tarjeta", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Tarjeta creada correctamente", Data: tarjeta}
}

func ObtenerTarjetaRFIDPorID(id uint) Response {
	tarjeta, err := db.ObtenerTarjetaRFIDPorID(id)
	if err != nil {
		return Response{Codigo: 400, Mensaje: "Error al obtener la tarjeta", Data: err.Error()}
	}

	return Response{Codigo: 200, Mensaje: "Tarjeta obtenida correctamente", Data: tarjeta}
}
