package db

func CreateSession(userID int, token string) error {

	query := `
	INSERT INTO sessions(user_id, token, expires_at)
	VALUES (?, ?, datetime('now', '+24 hours'))
	`

	_, err := DB.Exec(query, userID, token)
	return err
}
