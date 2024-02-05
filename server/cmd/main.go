package main

import (
	"context"
	"log/slog"
	"net/http"
	"pulsepathapi/bootstrap"
	"pulsepathapi/db"

	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := *app.EnvVariables
	database := app.Pconn
	err := database.Ping(context.Background())
	if err != nil {
		slog.Debug(err.Error())
	}
	defer db.ClosePsqlDatabase(database)

	router := gin.Default()

	router.GET("/ping", ping)

	router.Run(env["APP_URL"] + ":" + env["SERVER_PORT"])
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Ping is working"})
}
