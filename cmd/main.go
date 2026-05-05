package main

import (
	"forum/internal/db"
	"forum/internal/handlers"
	"net/http"
)

func main() {

	db.InitDB()

	db.CreateTables()

	http.HandleFunc("/register", handlers.Register)

	http.ListenAndServe(":8080", nil)

}
