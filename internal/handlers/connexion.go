package handlers

import (
	"net/http"

	"forum/internal/db"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/connexion.html")
		return
	}

	if r.Method == "POST" {

		identifiant := r.FormValue("identifiant")
		email := r.FormValue("email")
		mdp := r.FormValue("mdp")

		err := db.CreateUser(
			identifiant,
			email,
			mdp,
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.Write([]byte("user créé"))
	}
}
