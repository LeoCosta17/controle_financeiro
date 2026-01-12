package database

import (
	"database/sql"
	"log"
	"time"
)

func Init(url string, maxOpenConns, maxIdleConns int, maxIdleTime time.Duration) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "")
	if err != nil {
		log.Fatalf("error open db: %s", err)
	}

	db.SetMaxOpenConns(maxIdleConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
