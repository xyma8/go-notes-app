package handler

import (
	"encoding/json"
	"log"

	models "github.com/xyma8/go-notes-app/pkg/models"
	storage "github.com/xyma8/go-notes-app/pkg/storage"
)

func GetNote(noteID string) ([]byte, error) {
	noteData, err := storage.GetNoteById(noteID)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(noteData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func AddNote(noteDto models.NoteDto) {
	log.Printf(noteDto.Note)
}
