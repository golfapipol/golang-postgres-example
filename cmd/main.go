package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}
	fmt.Println("connected")

	http.ListenAndServe(":3000", nil)
}
