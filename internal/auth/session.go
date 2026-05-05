package auth

import (
	"net/http"

	"b1_projet-JS_Forum/internal/db"
)

func SetSessionCookie(w http.ResponseWriter, token string) {

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	})
}

func GetUserIDFromCookie(r *http.Request) (int, error) {

	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}

	var userID int

	err = db.DB.QueryRow(
		"SELECT user_id FROM sessions WHERE token = ?",
		cookie.Value,
	).Scan(&userID)

	return userID, err
}
