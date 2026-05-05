package db

import (
	"b1_projet-JS_Forum/internal/models"
)

func CreateUser(identifiant, email, mdp string) error {

	query := `
	INSERT INTO users(identifiant, email, mdp_hash)
	VALUES (?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		identifiant,
		email,
		mdp,
	)

	return err
}

func GetUserByEmail(email string) (models.User, error) {

	var user models.User

	query := `
	SELECT id, identifiant, email, mdp_hash
	FROM users
	WHERE email = ?
	`

	err := DB.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Identifiant,
		&user.Email,
		&user.Mdp,
	)

	return user, err
}

func DeleteUserByID(id int) error {

	query := `
	DELETE FROM users
	WHERE id = ?
	`

	_, err := DB.Exec(query, id)

	return err
}
