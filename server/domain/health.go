package domain

import "context"

type HealthRepository interface {
	CheckHealth(c context.Context) error
}

type HealthUsecase interface {
	CheckHealth(c context.Context) error
}
