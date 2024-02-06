package repository

import (
	"context"
	"log/slog"
	"pulsepathapi/domain"

	"github.com/jackc/pgx/v5"
)

type healthRepository struct {
	database *pgx.Conn
}

func NewHealthRepository(db *pgx.Conn) domain.HealthRepository {
	return &healthRepository{
		database: db,
	}
}

func (hr *healthRepository) CheckHealth(c context.Context) error {
	err := hr.database.Ping(c)
	if err != nil {
		slog.Info(err.Error())
	}
	return err
}
