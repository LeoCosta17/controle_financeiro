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
		url:          "/home/leonardo-costa/controle_financeiro/internal/database/db",
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
		log.Panic(err)
	}
	defer db.Close()

	repository := repositories.NewRepository(db)

	application := application{
		config:     config,
		repository: repository,
	}

	log.Fatal(application.run())
}
