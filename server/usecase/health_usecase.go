package usecase

import (
	"context"
	"pulsepathapi/domain"
	"time"
)

type healthUsecase struct {
	healthRepository domain.HealthRepository
	contextTimeout   time.Duration
}

func NewHealthUsecase(healthRepository domain.HealthRepository, timeout time.Duration) domain.HealthUsecase {
	return &healthUsecase{
		healthRepository: healthRepository,
		contextTimeout:   timeout,
	}
}

func (hu *healthUsecase) CheckHealth(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, hu.contextTimeout)
	defer cancel()
	return hu.healthRepository.CheckHealth(ctx)
}
