package db

func CreateSession(userID int, token string) error {

	query := `
	INSERT INTO sessions(user_id, token, expires_at)
	VALUES (?, ?, datetime('now', '+24 hours'))
	`

	_, err := DB.Exec(query, userID, token)
	return err
}

func GetUserIDFromToken(token string) (int, error) {

	var userID int

	err := DB.QueryRow(
		"SELECT user_id FROM sessions WHERE token = ? AND expires_at > datetime('now')",
		token,
	).Scan(&userID)

	return userID, err
}

func DeleteSession(token string) error {
	_, err := DB.Exec(
		"DELETE FROM sessions WHERE token = ?",
		token,
	)

	return err
}
