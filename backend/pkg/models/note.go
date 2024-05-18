package models

import (
	"time"
)

type Note struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	NoteText     string    `json:"note_text"`
	AuthorName   string    `json:"author_name"`
	CreationDate time.Time `json:"created_time"`
	IsPublic     bool      `json:"is_public"`
}
