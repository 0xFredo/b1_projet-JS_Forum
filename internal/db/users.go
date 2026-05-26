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

func GetUserRole(userID int) (string, error) {

	var role string

	query := `
	SELECT role
	FROM users
	WHERE id = ?
	`

	err := DB.QueryRow(
		query,
		userID,
	).Scan(&role)

	return role, err
}

func GetAllUsers() ([]models.User, error) {

	query := `
	SELECT id, identifiant, email, role
	FROM users
	ORDER BY id DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var u models.User

		err := rows.Scan(
			&u.ID,
			&u.Identifiant,
			&u.Email,
			&u.Role,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func UpdateUserRole(
	userID int,
	role string,
) error {

	query := `
	UPDATE users
	SET role = ?
	WHERE id = ?
	`

	_, err := DB.Exec(
		query,
		role,
		userID,
	)

	return err
}
