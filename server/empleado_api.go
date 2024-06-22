package server

import (
	"net/http"
	"strconv"

	"github.com/IsaiDiaz/bl"
	"github.com/IsaiDiaz/db"
	"github.com/gin-gonic/gin"
)

func registrarEmpleadoApi(router *gin.Engine) {
	router.POST("/empleado", crearEmpleado)
	router.GET("/empleado/:empid", obtenerEmpleadoPorID)
	router.POST("/empleado/:empid/tarjeta/:tarid", asignarTarjetaEmpleado)
	router.GET("/empleado/:empid/tarjeta", obtenerTarjetaEmpleado)
}

func crearEmpleado(c *gin.Context) {
	var empleado db.Empleado
	if err := c.ShouldBindJSON(&empleado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := bl.CrearEmpleado(&empleado)
	if response.Codigo == 400 {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func obtenerEmpleadoPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("empid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	response := bl.ObtenerEmpleadoPorID(uint(id))

	if response.Codigo == 400 {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func asignarTarjetaEmpleado(c *gin.Context) {
	empid, err := strconv.Atoi(c.Param("empid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleado inválido"})
		return
	}

	tarid, err := strconv.Atoi(c.Param("tarid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de tarjeta inválido"})
		return
	}

	empleadoTarjeta := db.EmpleadoTarjeta{
		EmpleadoID: uint(empid),
		TarjetaID:  uint(tarid),
	}

	response := bl.AsignarTarjetaEmpleado(&empleadoTarjeta)

	if response.Codigo == 400 {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func obtenerTarjetaEmpleado(c *gin.Context) {
	empid, err := strconv.Atoi(c.Param("empid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleado inválido"})
		return
	}

	response := bl.ObtenerTarjetaEmpleado(uint(empid))

	if response.Codigo == 400 {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
