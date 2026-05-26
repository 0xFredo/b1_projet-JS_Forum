package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"b1_projet-JS_Forum/internal/db"
)

func ViewPost(
	w http.ResponseWriter,
	r *http.Request,
) {

	idStr := r.URL.Query().Get("id")

	postID, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	post, err := db.GetPostByID(postID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	comments, err := db.GetCommentsByPostID(postID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := struct {
		Post     interface{}
		Comments interface{}
	}{
		Post:     post,
		Comments: comments,
	}

	tmpl := template.Must(
		template.ParseFiles(
			"web/templates/post.html",
		),
	)

	tmpl.Execute(w, data)
}

func CreateComment(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	cookie, err := r.Cookie("session_token")

	if err != nil {
		http.Error(w, "Not logged in", 401)
		return
	}

	userID, err := db.GetUserIDFromToken(
		cookie.Value,
	)

	if err != nil {
		http.Error(w, "Invalid session", 401)
		return
	}

	postID, _ := strconv.Atoi(
		r.FormValue("post_id"),
	)

	content := r.FormValue("content")

	err = db.CreateComment(
		postID,
		userID,
		content,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(
		w,
		r,
		"/post?id="+strconv.Itoa(postID),
		http.StatusSeeOther,
	)
}
