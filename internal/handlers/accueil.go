package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"b1_projet-JS_Forum/internal/db"
	"b1_projet-JS_Forum/internal/models"
)

func Home(w http.ResponseWriter, r *http.Request) {

	cat := strings.ToLower(r.URL.Query().Get("cat"))

	log.Println("FILTER =", cat)

	var posts []models.Post
	var err error

	if cat == "" {
		posts, err = db.GetAllPosts()
	} else {
		posts, err = db.GetPostsByCategory(cat)
	}

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	tmpl.Execute(w, posts)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/profile.html")
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/send_message.html")
}

func Photos(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/photos.html")
}

func Soirees(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/soirees.html")
}
