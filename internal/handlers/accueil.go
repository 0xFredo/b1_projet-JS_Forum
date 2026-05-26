package handlers

import (
	"html/template"
	"net/http"

	"b1_projet-JS_Forum/internal/db"
)

func Home(w http.ResponseWriter, r *http.Request) {

	posts, err := db.GetPostsByCategory("general")
	if err != nil {
		ErrorAlert(w, err.Error(), 500)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	tmpl.Execute(w, posts)
}

func Depeches(w http.ResponseWriter, r *http.Request) {

	posts, err := db.GetPostsByCategory("depeches")
	if err != nil {
		ErrorAlert(w, err.Error(), 500)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/depeches.html"))
	tmpl.Execute(w, posts)
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
