package db

import (
	"b1_projet-JS_Forum/internal/models"
)

func GetAllPosts() ([]models.Post, error) {

	var posts []models.Post

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

	ORDER BY posts.created_at DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var post models.Post

		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.ImagePath,
			&post.CreatedAt,
			&post.AuthorName,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
