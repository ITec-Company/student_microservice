package repositories

import (
	"database/sql"
	"student_microservice/internal/logging"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go

type App interface {
}

type Repository struct {
	App
}

func NewRepository(db *sql.DB, logger logging.Logger) *Repository {
	return &Repository{}
}
