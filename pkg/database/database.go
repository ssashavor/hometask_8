package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectionToPostgre() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=sasha dbname=contact")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	fmt.Println("Successfully connected!")
	return db, nil
}
