package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnect() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATASOURCE_URL"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
