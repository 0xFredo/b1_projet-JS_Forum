package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"b1_projet-JS_Forum/internal/db"

	"github.com/google/uuid"
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
		filePath := ""

		file, header, err := r.FormFile("piece_jointe")
		if err == nil {
			defer file.Close()

			filename := uuid.New().String() + filepath.Ext(header.Filename)
			filePath = "uploads/" + filename

			dst, err := os.Create(filePath)
			if err != nil {
				ErrorAlert(w, err.Error(), 500)
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				ErrorAlert(w, err.Error(), 500)
				return
			}
		}

		err = db.CreateMessage(
			userID,
			sujet,
			contenu,
			filePath,
		)

		if err != nil {
			ErrorAlert(w, err.Error(), 500)
			return
		}

		SuccessAlert(w, "Message envoyé", "/")
	}
}
