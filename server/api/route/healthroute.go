package route

import (
	"pulsepathapi/api/controller"
	"pulsepathapi/repository"
	"pulsepathapi/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewHealthRouter(env map[string]string, timeout time.Duration, db *pgx.Conn, group *gin.RouterGroup) {
	hr := repository.NewHealthRepository(db)
	hc := &controller.HealthController{
		HealthUsecase: usecase.NewHealthUsecase(hr, timeout),
	}
	group.GET("/health", hc.CheckHealth)
}
