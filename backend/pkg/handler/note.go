package handler

import (
	"encoding/json"
	"log"

	models "github.com/xyma8/go-notes-app/pkg/models"
	storage "github.com/xyma8/go-notes-app/pkg/storage"
)

func GetNote(noteId string) ([]byte, error) {
	noteData, err := storage.GetNoteById(noteId)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(noteData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func AddNote(noteDto models.NoteDto) ([]byte, error) {
	log.Printf(noteDto.Note)
	noteUUID, err := storage.AddNewNote(noteDto)
	if err != nil {
		log.Printf("Ошибка при добавлении новой заметки: %v", err)
		return nil, err
	}

	jsonData, err := json.Marshal(noteUUID)
	if err != nil {
		return nil, err
	}

	return jsonData, nil

}
