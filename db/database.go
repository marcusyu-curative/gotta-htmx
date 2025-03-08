package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type Database struct {
    Conn *sql.DB
}

func New() (*Database, error) {
    conn, err := sql.Open("sqlite3", "./awesome.db")
	if err != nil {
		log.Fatal(err)
        return nil, err
	}

    db := &Database{Conn: conn}
    
	if err := db.runMigrations(); err != nil {
		log.Fatal(err)
        return nil, err
	}
    
    return db, nil
}

func (db *Database) runMigrations() error {
    _, err := db.Conn.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		status TEXT
	);`)
    if err != nil {
        log.Println("Error running migrations:", err)
        return err
    }
    
    log.Println("Database migrations applied successfully.")
    return nil
}

func (db *Database) Close() {
    if db.Conn != nil {
        db.Conn.Close()
    }
}
