package router

import (
	"net/http"
    "encoding/json"
	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	Status string `json:"status"`
}

// Serialize Healthcheck response to string
func (h *Healthcheck)String() string {
    serialized_response, _ := json.Marshal(h)

    return string(serialized_response)
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
