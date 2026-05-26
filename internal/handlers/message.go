package handlers

import (
	"net/http"

	"b1_projet-JS_Forum/internal/db"
)

func SendMessageAdmin(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method == "GET" {

		http.ServeFile(
			w,
			r,
			"web/templates/send_message.html",
		)

		return
	}

	if r.Method == "POST" {

		cookie, err := r.Cookie("session_token")
		if err != nil {
			ErrorAlert(w, "Non connecté", 401)
			return
		}

		userID, err := db.GetUserIDFromToken(cookie.Value)
		if err != nil {
			ErrorAlert(w, "Session invalide", 401)
			return
		}

		sujet := r.FormValue("sujet")
		contenu := r.FormValue("contenu")

		err = db.CreateMessage(
			userID,
			sujet,
			contenu,
		)

		if err != nil {
			ErrorAlert(w, err.Error(), 500)
			return
		}

		w.Write([]byte("Message envoyé"))
	}
}
