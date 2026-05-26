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
	http.HandleFunc("/create-post", handlers.CreatePost)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/post", handlers.ViewPost)
	http.HandleFunc("/create-comment", handlers.CreateComment)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/all-post", handlers.Home)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
