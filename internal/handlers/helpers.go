package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"b1_projet-JS_Forum/internal/db"
)

// GetUserIDFromCookie récupère l'ID de l'utilisateur connecté depuis le cookie
func GetUserIDFromCookie(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}
	return db.GetUserIDFromToken(cookie.Value)
}

// GetUserRoleFromCookie récupère le rôle de l'utilisateur connecté
func GetUserRoleFromCookie(r *http.Request) (string, error) {
	userID, err := GetUserIDFromCookie(r)
	if err != nil {
		return "", err
	}
	return db.GetUserRole(userID)
}

// CheckPermission vérifie que l'utilisateur a le bon rôle
// Retourne true si autorisé, false sinon
func CheckPermission(r *http.Request, requiredRole string) bool {
	role, err := GetUserRoleFromCookie(r)
	if err != nil {
		return false
	}
	return role == requiredRole
}

// CheckPermissionMultiple vérifie que l'utilisateur a l'un des rôles autorisés
func CheckPermissionMultiple(r *http.Request, allowedRoles ...string) bool {
	role, err := GetUserRoleFromCookie(r)
	if err != nil {
		return false
	}
	for _, allowed := range allowedRoles {
		if role == allowed {
			return true
		}
	}
	return false
}

func NoCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

func ErrorAlert(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	escapedMsg, _ := json.Marshal(message)
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<script>
		alert(%s);
		window.history.back();
	</script>
</head>
<body></body>
</html>`, string(escapedMsg))
}
func SuccessAlert(w http.ResponseWriter, message string, redirectURL string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)

	escapedMsg, _ := json.Marshal(message)
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<script>
		alert(%s);
		window.location.href = %s;
	</script>
</head>
<body></body>
</html>`, string(escapedMsg), fmt.Sprintf(`"%s"`, redirectURL))
}
