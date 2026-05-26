package models

import "time"

type Message struct {
	ID           int
	UserID       int
	AuthorName   string
	Sujet        string
	Contenu      string
	DateCreation time.Time
}
