package db

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
