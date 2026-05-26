package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"b1_projet-JS_Forum/internal/db"
	"b1_projet-JS_Forum/internal/models"
)

func AdminPage(
	w http.ResponseWriter,
	r *http.Request,
) {

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Non connecté", 401)
		return
	}

	userID, err := db.GetUserIDFromToken(cookie.Value)
	if err != nil {
		http.Error(w, "Session invalide", 401)
		return
	}

	role, err := db.GetUserRole(userID)
	if err != nil {
		http.Error(w, "Erreur role", 500)
		return
	}

	if role != "admin" {
		http.Error(w, "Forbidden", 403)
		return
	}

	messages, err := db.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := struct {
		Messages []models.Message
		Users    []models.User
	}{
		Messages: messages,
		Users:    users,
	}

	tmpl := template.Must(
		template.ParseFiles(
			"web/templates/admin/inbox.html",
		),
	)

	tmpl.Execute(w, data)
}

func PromoteUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	idStr := r.FormValue("user_id")

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	err = db.UpdateUserRole(
		userID,
		"commere",
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(
		w,
		r,
		"/admin",
		http.StatusSeeOther,
	)
}
