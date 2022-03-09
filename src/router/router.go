package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	Status string `json:"status"`
}

// PingExample godoc
// @Summary Healthcheck
// @Schemes
// @Description Liveness Probe
// @Tags Default
// @Accept json
// @Produce json
// @Success 200 {object} router.Healthcheck
// @Router /healthcheck [get]
func healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func InitRoutes(baseGroup *gin.RouterGroup) {
	baseGroup.GET("/healthcheck", healthcheck)
}
