package infra

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	DB *sql.DB
}

func Setup() (*DB, error) {
	log.Println("Fake DB")
	return nil, nil
}
