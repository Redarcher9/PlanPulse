package main

import (
	"context"
	"log/slog"
	"pulsepathapi/api/route"
	"pulsepathapi/bootstrap"
	"pulsepathapi/db"
	"time"

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
	route.SetupRoutes(env, 3*time.Second, database, router)
	router.Run(env["APP_URL"] + ":" + env["SERVER_PORT"])
}
