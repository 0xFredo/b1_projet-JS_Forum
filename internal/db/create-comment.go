package db

import "b1_projet-JS_Forum/internal/models"

func CreateComment(
	postID int,
	userID int,
	content string,
) error {

	query := `
	INSERT INTO comments(
		post_id,
		user_id,
		content
	)
	VALUES (?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		postID,
		userID,
		content,
	)

	return err
}

func GetCommentsByPostID(
	postID int,
) ([]models.Comment, error) {

	var comments []models.Comment

	query := `
	SELECT
		comments.id,
		comments.post_id,
		comments.user_id,
		comments.content,
		comments.created_at,
		users.identifiant

	FROM comments

	INNER JOIN users
	ON comments.user_id = users.id

	WHERE comments.post_id = ?

	ORDER BY comments.created_at ASC
	`

	rows, err := DB.Query(query, postID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var comment models.Comment

		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.AuthorName,
		)

		if err != nil {
			return nil, err
		}

		comments = append(
			comments,
			comment,
		)
	}

	return comments, nil
}
