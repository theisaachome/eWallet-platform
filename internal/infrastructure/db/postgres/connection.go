package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"time"
)

func NewPostgresDB() *sqlx.DB {
	dbUsr := os.Getenv("DB_USR")
	dbPassWd := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbNAME := os.Getenv("DB_NAME")

	// PostgreSQL DSN format: postgres://username:password@host:port/dbname?sslmode=disable
	//dataSource := fmt.Sprintf(
	//	"postgres://%s:%s@%s:%s/%s?sslmode=disable",
	//	dbUsr, dbPassWd, dbHost, dbPORT, dbNAME,
	//)
	//dataSource := fmt.Sprintf(
	//	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	dbHost, dbPORT, dbUsr, dbPassWd, "ewallet_db",
	//)
	dataSource := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPORT, dbUsr, dbPassWd, dbNAME,
	)

	db, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		panic(fmt.Errorf("could not open db: %w", err))
	}

	db.Ping()
	row := db.QueryRow("SELECT current_database();")
	var dbName string
	row.Scan(&dbName)
	fmt.Println("Connected to DB:", dbName)

	if err := db.Ping(); err != nil {

		panic(fmt.Errorf("could not ping db: %w", err))
	} else {
		fmt.Println("✅ Successfully connected to PostgreSQL")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
