package main

import (
	"app/internal/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config     config
	repository repositories.Repository
}

type config struct {
	api_port string
	db       dbConfig
	router   chi.Mux
}

type dbConfig struct {
	url          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func (app *application) run() error {

	server := &http.Server{
		Addr:         app.config.api_port,
		Handler:      &app.config.router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at: %s", app.config.api_port)

	return server.ListenAndServe()
}
