package handlers

import (
	"html/template"
	"net/http"

	"b1_projet-JS_Forum/internal/db"
)

func Home(w http.ResponseWriter, r *http.Request) {

	posts, err := db.GetAllPosts()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl := template.Must(
		template.ParseFiles(
			"web/templates/index.html",
		),
	)

	tmpl.Execute(
		w,
		posts,
	)
}
