package bootstrap

import (
	"pulsepathapi/db"

	"github.com/jackc/pgx/v5"
)

type AppInstance struct {
	EnvVariables *map[string]string
	Pconn        *pgx.Conn
}

func App() *AppInstance {
	app := &AppInstance{}
	app.EnvVariables = NewEnv()
	app.Pconn = db.NewPsqlDatabase(*app.EnvVariables)
	return app
}
