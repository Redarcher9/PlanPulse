package db

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

/*
Description: Wrapper struct for Postgresql Database
*/
type PsqlDatabase struct {
	dbInstance *pgx.Conn
}

/*
	Description: Build Connection String for connecting database

	Variables: Environment Variables

	Returns: Connection String
*/

func BuildConnString(EnvVariables map[string]string) string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	connString := "postgres://" + EnvVariables["DATABASE_USERNAME"] + ":" + EnvVariables["DATABASE_PASSWORD"] + "@" + EnvVariables["DATABASE_ADDRESS"] + ":" + EnvVariables["SERVER_PORT"] + "/" + EnvVariables["DATABASE_NAME"]
	return connString
}

/*
	Description: Connect to PostGresql Database and Return Wrapper Database Object

	Variables: Environment Variables

	Returns: PsqlDatabase Wrapper Object, error if any
*/

func NewPsqlDatabase(EnvVariables map[string]string) (*PsqlDatabase, error) {
	conn, err := pgx.Connect(context.Background(), BuildConnString(EnvVariables))
	if err != nil {
		slog.Warn(err.Error())
		return &PsqlDatabase{dbInstance: conn}, err
	}
	return &PsqlDatabase{dbInstance: conn}, nil
}
