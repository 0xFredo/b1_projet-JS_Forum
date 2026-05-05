package main

import (
	"log"
	"net/http"

	"b1_projet-JS_Forum/internal/db"
	"b1_projet-JS_Forum/internal/handlers"
)

func main() {

	db.InitDB()
	db.CreateTables()

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/delete-user", handlers.DeleteUser)
	//http.HandleFunc("/create-post", auth.RequireAuth(handlers.CreatePost))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
