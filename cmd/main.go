package main

import (
	"b1_projet-JS_Forum/internal/db"
	"b1_projet-JS_Forum/internal/handlers"
	"net/http"
)

func main() {

	db.InitDB()

	db.CreateTables()

	http.HandleFunc("/register", handlers.Register)

	http.ListenAndServe(":8080", nil)

}
