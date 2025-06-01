package web

import "github.com/gin-gonic/gin"

type Routable interface {
	RegisterRoutes(router *gin.RouterGroup)
}
