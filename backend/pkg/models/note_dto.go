package models

type NoteDto struct {
	Title  string `json:"title,omitempty"`
	Note   string `json:"note"`
	Author string `json:"author,omitempty"`
}
