package main

import (
	"context"
	"log"
	"os/signal"
	"student_microservice/config"
	"student_microservice/handlers"
	"student_microservice/internal/databases"
	"student_microservice/internal/logging"
	"student_microservice/internal/server"
	"student_microservice/repositories"
	"student_microservice/services"
	"syscall"
	"time"
)

func main() {
	logger := logging.GetLoggerLogrus()
	conf := config.GetConfig(logger)

	db, err := postgres.NewPostgresDB(&postgres.PostgresDB{
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		Username: conf.DB.Username,
		Password: conf.DB.Password,
		DBName:   conf.DB.DBName,
		SSLMode:  conf.DB.SSLMode,
		Logger:   logger,
	})
	if err != nil {
		log.Panicf("Error while initialization database:%s", err)
	}

	repo := repositories.NewRepository(db, logger)

	ser := services.NewService(repo, logger)
	handler := handlers.NewHandler(ser, logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	serv := new(server.Server)
	logger.Infof("Starting server on %s:%s...", conf.Server.Host, conf.Server.Port)
	go func() {
		if err := serv.Run(conf.Server.Host, conf.Server.Port, handler.InitRoutes()); err != nil {
			logger.Panicf("Error occured while running http server: %s", err.Error())
		}
	}()

	<-ctx.Done()
	stop()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		if err = db.Close(); err != nil {
			logger.Errorf("Error while closing database:%s", err)
		}
		logger.Info("Database closed")
		cancel()
	}()
	if err = serv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown:%s", err)
	}
}
