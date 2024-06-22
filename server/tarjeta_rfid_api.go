package server

import (
	"net/http"
	"strconv"

	"github.com/IsaiDiaz/bl"
	"github.com/IsaiDiaz/db"
	"github.com/gin-gonic/gin"
)

func registrarTarjetaApi(router *gin.Engine) {
	router.POST("/tarjeta", crearTarjetaRFID)
	router.GET("/tarjeta/:id", obtenerTarjetaRFIDPorID)
	router.GET("/tarjeta", obtenerTodasTarjetasRFID)
}

func crearTarjetaRFID(c *gin.Context) {
	var tarjeta db.Tarjeta
	if err := c.ShouldBindJSON(&tarjeta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := bl.CrearTarjetaRFID(&tarjeta)
	if response.Codigo == 400 {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func obtenerTarjetaRFIDPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	response := bl.ObtenerTarjetaRFIDPorID(uint(id))

	if response.Codigo == 400 {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func obtenerTodasTarjetasRFID(c *gin.Context) {
	response := bl.ObtenerTodasTarjetasRFID()

	if response.Codigo == 400 {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
