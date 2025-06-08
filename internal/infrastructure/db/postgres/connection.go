package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func NewPostgresDB() *sql.DB {
	dbUsr := os.Getenv("DB_USR")
	dbPassWd := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbNAME := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsr, dbPassWd, dbHost, dbPORT, dbNAME)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	// Check if the database is reachable
	if err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
