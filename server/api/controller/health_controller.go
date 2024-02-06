package controller

import (
	"log/slog"
	"net/http"
	"pulsepathapi/domain"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	HealthUsecase domain.HealthUsecase
}

func (pc *HealthController) CheckHealth(c *gin.Context) {
	err := pc.HealthUsecase.CheckHealth(c)
	if err != nil {
		slog.Info(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Server and Database are running!",
	})
	return
}
