package handlers

import (
	"html/template"
	"net/http"

	"b1_projet-JS_Forum/internal/db"
)

func Home(
	w http.ResponseWriter,
	r *http.Request,
) {

	posts, err := db.GetAllPosts()

	if err != nil {
		http.Error(
			w,
			err.Error(),
			500,
		)
		return
	}

	tmpl := template.Must(
		template.ParseFiles("web/templates/index.html"),
	)

	tmpl.Execute(
		w,
		posts,
	)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/profile.html")
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/send_message.html")
}

func Archive(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/archive.html")
}

func Photos(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/photos.html")
}

func Soirees(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/soirees.html")
}
