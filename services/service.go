package services

import (
	"student_microservice/internal/logging"
	"student_microservice/repositories"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type App interface {
}

type Service struct {
	App
}

func NewService(repository *repositories.Repository, logger logging.Logger) *Service {
	return &Service{}
}
