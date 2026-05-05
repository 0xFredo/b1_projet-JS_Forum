package handlers

import (
	"net/http"
	"strconv"

	"b1_projet-JS_Forum/internal/db"

	"golang.org/x/crypto/bcrypt"
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

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/login.html")
		return
	}

	if r.Method == "POST" {

		email := r.FormValue("email")
		mdp := r.FormValue("mdp")

		user, err := db.GetUserByEmail(email)

		if err != nil {
			http.Error(
				w,
				"Utilisateur introuvable",
				http.StatusUnauthorized,
			)
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(user.Mdp),
			[]byte(mdp),
		)

		if err != nil {
			http.Error(
				w,
				"Mot de passe incorrect",
				http.StatusUnauthorized,
			)
			return
		}

		w.Write([]byte("Connexion réussie"))
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")

	if idStr == "" {
		http.Error(w, "id manquant", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	err = db.DeleteUserByID(id)
	if err != nil {
		http.Error(w, "erreur suppression user", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("user supprimé"))
}
