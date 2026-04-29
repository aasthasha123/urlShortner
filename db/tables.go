package db

import (
	"fmt"
	"log"
)

func CreateTables() {
	db := SetDB()
	fmt.Println("GOT DB CONNECTION")
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		longurl TEXT NOT NULL,
		shorturl TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	h, err := db.Exec(query)
	fmt.Println("CREATED : ", h, err)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
