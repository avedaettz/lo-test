package app

import (
	"context"
	"log"
	"lotest/internal/handler"
	"lotest/internal/logger"
	"lotest/internal/repository"
	"lotest/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server *http.Server
	logger *logger.Logger
}

func NewApp() *App {
	repo := repository.NewRepo()
	logger := logger.NewLogger()
	taskService := service.NewTaskService(repo)
	taskHandler := handler.NewTaskHandler(taskService, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", taskHandler.GetTasks)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetTask)
	mux.HandleFunc("POST /tasks", taskHandler.CreateTask)

	return &App{
		server: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
		logger: logger,
	}
}

func (a *App) Run() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	a.logger.Close()
}
