package auth

import (
	"b1_projet-JS_Forum/internal/db"
	"net/http"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Not logged in", 401)
			return
		}

		var userID int
		err = db.DB.QueryRow(
			"SELECT user_id FROM sessions WHERE token = ?",
			cookie.Value,
		).Scan(&userID)

		if err != nil {
			http.Error(w, "Session invalide", 401)
			return
		}

		next(w, r)
	}
}
