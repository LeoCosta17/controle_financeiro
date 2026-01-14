package main

import (
	"app/internal/database"
	"app/internal/repositories"
	"app/internal/services"
	"log"
	"time"
)

func main() {

	dbConfig := dbConfig{
		url:          "/home/leonardo-costa/Documentos/producao.db",
		maxOpenConns: 10,
		maxIdleConns: 4,
		maxIdleTime:  time.Second * 20,
	}

	config := config{
		api_port: ":5000",
		db:       dbConfig,
	}

	db, err := database.Init(
		config.db.url,
		config.db.maxOpenConns,
		config.db.maxIdleConns,
		config.db.maxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repositories.NewRepository(db)
	services := services.NewServices(repository)
	r := LoadRoutes(services)

	application := application{
		config:     config,
		repository: repository,
		services:   services,
	}

	if err := application.initTables(db); err != nil {
		log.Fatal(err)
	}

	if err := application.run(r); err != nil {
		log.Fatal(err)
	}

	log.Printf("server has started port: %s", config.api_port)
}
