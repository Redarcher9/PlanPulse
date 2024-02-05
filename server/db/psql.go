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
	connString := "postgres://" + EnvVariables["DATABASE_USERNAME"] + ":" + EnvVariables["DATABASE_PASSWORD"] + "@" + EnvVariables["DATABASE_ADDRESS"] + ":" + EnvVariables["DATABASE_PORT"] + "/" + EnvVariables["DATABASE_NAME"]
	println(connString)
	return connString
}

/*
	Description: Connect to PostGresql Database and Return Wrapper Database Object

	Variables: Environment Variables

	Returns: PsqlDatabase Wrapper Object, error if any
*/

func NewPsqlDatabase(EnvVariables map[string]string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), BuildConnString(EnvVariables))
	if err != nil {
		slog.Warn(err.Error())
		return conn
	}
	return conn
}

/*
Description: Close the connection to PostGresql Database

Variables: Database connection

Returns: NA
*/
func ClosePsqlDatabase(dbconn *pgx.Conn) {
	dbconn.Close(context.Background())
}
