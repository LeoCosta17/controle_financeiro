package main

import (
	"app/internal/repositories"
	"app/internal/services"
	"database/sql"
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

func (app *application) initTables(db *sql.DB) error {

	query := `
		CREATE TABLE IF NOT EXISTS "users" (
		"id"	INTEGER,
		"name"	TEXT NOT NULL,
		"email"	TEXT NOT NULL UNIQUE,
		"password"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	query = `
		CREATE TABLE IF NOT EXISTS "suppliers" (
		"id"	INTEGER,
		"name"	TEXT NOT NULL UNIQUE,
		"email"	TEXT UNIQUE,
		"telephone"	TEXT UNIQUE,
		"document" TEXT UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)
	)
	`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	query = `
		CREATE TABLE IF NOT EXISTS "costumers"(
		"id" INTEGER,
		"name" TEXT NOT NULL UNIQUE,
		"email" TEXT UNIQUE,
		"telephone" TEXT UNIQUE,
		"document" TEXT UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)	
		)
	`

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}
