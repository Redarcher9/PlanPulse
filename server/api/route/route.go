package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

/*
@Description: Sets up Routes for API
@Parameters: Environment variables, timeout seconds, database Instance
*/
func SetupRoutes(env map[string]string, timeout time.Duration, db *pgx.Conn, gin *gin.Engine) {
	Router := gin.Group("")

	//All APIs are declared here
	NewHealthRouter(env, timeout, db, Router)
}
