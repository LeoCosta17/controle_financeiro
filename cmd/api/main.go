package main

import (
	"app/internal/database"
	"app/internal/repositories"
	"log"
	"time"
)

func main() {

	r := LoadRoutes()

	dbConfig := dbConfig{
		url:          "/home/leonardo-costa/controle_financeiro/internal/database/producao.db",
		maxOpenConns: 10,
		maxIdleConns: 4,
		maxIdleTime:  time.Second * 20,
	}

	config := config{
		api_port: ":5000",
		db:       dbConfig,
		router:   *r,
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

	application := application{
		config:     config,
		repository: repository,
	}

	if err := application.run(); err != nil {
		log.Fatal(err)
	}

	log.Printf("server has started port: %s", config.api_port)
}
