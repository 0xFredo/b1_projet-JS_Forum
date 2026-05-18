package db

func CreatePost(userID int, title, content, imagePath string) error {

	query := `
	INSERT INTO posts(user_id, title, content, image_path)
	VALUES (?, ?, ?, ?)
	`

	_, err := DB.Exec(query, userID, title, content, imagePath)
	return err
}
