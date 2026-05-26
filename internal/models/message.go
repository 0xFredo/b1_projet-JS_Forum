package models

import (
	"strings"
	"time"
)

type Message struct {
	ID           int
	UserID       int
	AuthorName   string
	Sujet        string
	Contenu      string
	FilePath     string
	DateCreation time.Time
}

func (m Message) HasFile() bool {
	return m.FilePath != ""
}

func (m Message) IsImageFile() bool {
	filePath := strings.ToLower(m.FilePath)

	return strings.HasSuffix(filePath, ".jpg") ||
		strings.HasSuffix(filePath, ".jpeg") ||
		strings.HasSuffix(filePath, ".png") ||
		strings.HasSuffix(filePath, ".gif") ||
		strings.HasSuffix(filePath, ".webp")
}
