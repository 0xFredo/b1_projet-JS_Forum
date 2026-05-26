package db

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateAdminIfNotExists() {

	var count int

	query := `
	SELECT COUNT(*)
	FROM users
	WHERE role = 'admin'
	`

	err := DB.QueryRow(query).Scan(&count)

	if err != nil {
		log.Println(err)
		return
	}

	if count > 0 {
		log.Println("Admin déjà existant")
		return
	}

	password := "babylon"

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		log.Println(err)
		return
	}

	insertQuery := `
	INSERT INTO users(
		identifiant,
		email,
		mdp_hash,
		role
	)
	VALUES (?, ?, ?, ?)
	`

	_, err = DB.Exec(
		insertQuery,
		"gg",
		"gg@gossip.com",
		string(hash),
		"admin",
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return
		}

		log.Println(err)
		return
	}

	log.Println("Compte admin créé")
}
