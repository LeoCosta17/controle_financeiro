package main

import (
	"app/internal/repositories"
	"app/internal/services"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config     config
	repository repositories.Repository
	services   services.Services
}

type config struct {
	api_port string
	db       dbConfig
}

type dbConfig struct {
	url          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func (app *application) run(r *chi.Mux) error {

	server := &http.Server{
		Addr:         app.config.api_port,
		Handler:      r,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	fmt.Printf("server has started at: %s\n", app.config.api_port)

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
