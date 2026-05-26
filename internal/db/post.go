package db

import (
	"b1_projet-JS_Forum/internal/models"
)

func CreatePost(userID int, title, content, imagePath string) error {

	query := `
	INSERT INTO posts(user_id, title, content, image_path)
	VALUES (?, ?, ?, ?)
	`

	_, err := DB.Exec(query, userID, title, content, imagePath)
	return err
}

func GetPostByID(id int) (models.Post, error) {

	var post models.Post

	query := `
	SELECT
		posts.id,
		posts.user_id,
		posts.title,
		posts.content,
		posts.image_path,
		posts.created_at,
		users.identifiant

	FROM posts

	INNER JOIN users
	ON posts.user_id = users.id

	WHERE posts.id = ?
	`

	err := DB.QueryRow(
		query,
		id,
	).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.ImagePath,
		&post.CreatedAt,
		&post.AuthorName,
	)

	return post, err
}
