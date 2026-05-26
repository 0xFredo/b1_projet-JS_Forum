package db

import (
	"b1_projet-JS_Forum/internal/models"
)

func CreateUser(
	identifiant,
	email,
	mdp,
	role string,
) error {

	query := `
	INSERT INTO users(
		identifiant,
		email,
		mdp_hash,
		role
	)
	VALUES (?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		identifiant,
		email,
		mdp,
		role,
	)

	return err
}

func GetUserByEmail(email string) (models.User, error) {

	var user models.User

	query := `
	SELECT id, identifiant, email, mdp_hash, role
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
		&user.Role,
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
