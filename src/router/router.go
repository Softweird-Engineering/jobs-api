package router

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes ...
func InitRoutes(baseGroup *gin.RouterGroup) {
	baseGroup.GET("/healthcheck", healthcheck)
}
