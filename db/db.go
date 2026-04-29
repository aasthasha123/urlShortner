package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func SetDB() *sqlx.DB {
	//DB Connection

	db, err := sqlx.Connect("postgres", "user=myuser dbname=mydb sslmode=disable password=mypassword host=localhost")
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
	return db
}
