package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "./todo.db"
		log.Printf("Environment variable DATABASE_URL not set. Using default url: %s", dbUrl)
	} else {
		log.Printf("Using dbUrl from environment variable: %s", dbUrl)
	}

	var err error
	db, err = sql.Open("sqlite3", dbUrl)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	log.Print("Database initialized successfully")
	runMigrations()
}

func CloseDB() {
	if db != nil {
		db.Close()
		log.Print("Database connection closed")
	}
}

func GetDB() *sql.DB {
	if db == nil {
		log.Print("Database connection is not initialized")
		log.Fatal("Database connection is not initialized")
	}
	return db
}

func runMigrations() {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Print("Database migrations completed")
}
