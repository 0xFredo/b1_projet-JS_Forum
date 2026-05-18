package models

import "time"

type Post struct {
	ID         int
	UserID     int
	Title      string
	Content    string
	ImagePath  string
	CreatedAt  time.Time
	AuthorName string
}
