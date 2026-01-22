package infra

import (
	"database/sql"
	"log"
)

type DB struct {
	DB *sql.DB
}

func Setup() (*DB, error) {
	log.Println("Fake DB")
	return nil, nil
}

func (db *DB) HelloWorld() {
	log.Println("hello from db")
}
