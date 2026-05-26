package db

import (
	"log"
	"strings"
)

func CreateTables() {

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		identifiant TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		mdp_hash TEXT NOT NULL,
		role TEXT NOT NULL CHECK(role IN ('user', 'commere', 'admin')) DEFAULT 'user',
		date_création DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		token TEXT UNIQUE NOT NULL,
		expires_at DATETIME NOT NULL,

		FOREIGN KEY(user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nom TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	image_path TEXT,
	category_id INTEGER,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY(user_id) REFERENCES users(id),
	FOREIGN KEY(category_id) REFERENCES categories(id)
);

	CREATE TABLE IF NOT EXISTS post_categories (
		post_id INTEGER,
		category_id INTEGER,

		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(category_id) REFERENCES categories(id)
	);

	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,

		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,

		content TEXT NOT NULL,

		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS reactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		target_type TEXT NOT NULL,
		target_id INTEGER NOT NULL,
		valeur INTEGER NOT NULL,

		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    user_id INTEGER NOT NULL,

    sujet TEXT NOT NULL,

    contenu TEXT NOT NULL,

    file_path TEXT,

    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(user_id) REFERENCES users(id)
);
	`

	_, err := DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec("ALTER TABLE messages ADD COLUMN file_path TEXT")
	if err != nil && !strings.Contains(err.Error(), "duplicate column name") {
		log.Println("Migration messages.file_path:", err)
	}

	log.Println("Tables creé")
}
