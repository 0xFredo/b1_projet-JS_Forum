package handlers

import (
	"net/http"
	"strconv"

	"b1_projet-JS_Forum/internal/auth"
	"b1_projet-JS_Forum/internal/db"

	"github.com/google/uuid"

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

		var userID int
		var hash string

		err := db.DB.QueryRow(
			"SELECT id, mdp_hash FROM users WHERE email = ?",
			email,
		).Scan(&userID, &hash)

		if err != nil {
			http.Error(w, "User not found", 401)
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(hash),
			[]byte(mdp),
		)

		if err != nil {
			http.Error(w, "Mauvais mot de passe", 401)
			return
		}

		cookie := uuid.New().String()

		err = db.CreateSession(userID, cookie)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		auth.SetSessionCookie(w, cookie)

		w.Write([]byte("Connecté"))
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

func CreatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/create_post.html")
		return
	}

	if r.Method == "POST" {

		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Not logged in", 401)
			return
		}

		userID, err := db.GetUserIDFromToken(cookie.Value)
		if err != nil {
			http.Error(w, "Invalid session", 401)
			return
		}

		title := r.FormValue("title")
		content := r.FormValue("content")

		imagePath := ""

		err = db.CreatePost(userID, title, content, imagePath)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("Post créé"))
	}
}
