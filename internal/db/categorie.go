package db

import (
	"b1_projet-JS_Forum/internal/models"
	"log"
)

func SeedCategories() {

	categories := []string{
		"general",
		"depeches",
	}

	for _, c := range categories {

		_, err := DB.Exec(
			"INSERT OR IGNORE INTO categories (nom) VALUES (?)",
			c,
		)

		if err != nil {
			log.Println("Seed error:", err)
		}
	}

	generalID, err := GetCategoryIDByName("general")
	if err != nil {
		log.Println("Category cleanup error:", err)
		return
	}

	_, err = DB.Exec(`
		UPDATE posts
		SET category_id = ?
		WHERE category_id IN (
			SELECT id
			FROM categories
			WHERE nom IN ('photos', 'soirees')
		)
	`, generalID)

	if err != nil {
		log.Println("Category cleanup error:", err)
	}

	_, err = DB.Exec("DELETE FROM categories WHERE nom IN ('photos', 'soirees')")
	if err != nil {
		log.Println("Category cleanup error:", err)
	}

	log.Println("Categories seed OK")
}

func GetCategoryIDByName(name string) (int, error) {

	var id int

	err := DB.QueryRow(
		"SELECT id FROM categories WHERE nom = ?",
		name,
	).Scan(&id)

	return id, err
}

func GetPostsByCategory(cat string) ([]models.Post, error) {

	rows, err := DB.Query(`
		SELECT posts.id, posts.title, posts.content, posts.image_path,
		       posts.created_at,
		       users.identifiant
		FROM posts
		JOIN users ON users.id = posts.user_id
		JOIN categories ON categories.id = posts.category_id
		WHERE categories.nom = ?
		ORDER BY posts.created_at DESC
	`, cat)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post

		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.ImagePath,
			&p.CreatedAt,
			&p.AuthorName,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, nil
}
