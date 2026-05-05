package handlers

import (
	"net/http"

<<<<<<< HEAD
	"forum/internal/db"

	"golang.org/x/crypto/bcrypt"
=======
	"b1_projet-JS_Forum/internal/db"
>>>>>>> refs/remotes/origin/main
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

		mdp_hash, err := bcrypt.GenerateFromPassword(
			[]byte(mdp),
			bcrypt.DefaultCost,
		)

		if err != nil {
			http.Error(w, "Erreur hash", 500)
			return
		}

		err = db.CreateUser(
			identifiant,
			email,
			string(mdp_hash),
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("user créé"))
	}
}
