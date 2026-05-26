package db

import "b1_projet-JS_Forum/internal/models"

func CreateMessage(
	userID int,
	sujet string,
	contenu string,
) error {

	query := `
	INSERT INTO messages(
		user_id,
		sujet,
		contenu
	)
	VALUES (?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		userID,
		sujet,
		contenu,
	)

	return err
}

func GetAllMessages() ([]models.Message, error) {

	query := `
	SELECT
		messages.id,
		messages.sujet,
		messages.contenu,
		messages.date_creation,
		users.identifiant
	FROM messages
	INNER JOIN users
	ON messages.user_id = users.id
	ORDER BY messages.date_creation DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []models.Message

	for rows.Next() {

		var m models.Message

		err := rows.Scan(
			&m.ID,
			&m.Sujet,
			&m.Contenu,
			&m.DateCreation,
			&m.AuthorName,
		)

		if err != nil {
			return nil, err
		}

		messages = append(messages, m)
	}

	return messages, nil
}
