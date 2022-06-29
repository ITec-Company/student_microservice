package handlers

import (
	"github.com/gorilla/mux"
	"student_microservice/internal/logging"
	"student_microservice/services"
)

type Handler struct {
	service *services.Service
	logger  logging.Logger
}

func NewHandler(service *services.Service, logger logging.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	return r
}
