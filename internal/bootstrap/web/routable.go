package web

import "github.com/gin-gonic/gin"

type Routable interface {
	RegisterRoutes(router *gin.RouterGroup)
	// RegisterPublicRoutes(router *gin.RouterGroup)
	// RegisterProtectedRoutes(router *gin.RouterGroup)
}
