package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	registrarTarjetaApi(router)

	log.Fatal(router.Run(":8080"))

}
