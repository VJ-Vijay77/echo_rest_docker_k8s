package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	DSN := "postgresql://vijay:zmxmcmvbn@postgres:5432/echo_rest?sslmode=disable"
	db, err := sqlx.Connect("postgres", DSN)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
