package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	models "github.com/xyma8/go-notes-app/pkg/models"
)

var db *sql.DB

func ConnectDB(dsn string) {
	var err error
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка проверки соединения: %v", err)
	}

	log.Printf("Успешное подключение к базе данных!")
}

func GetNoteById(noteID string) (*models.Note, error) {
	query := "SELECT id, title, note_text, author, creation_time, is_public FROM notes WHERE id = ?"
	var note models.Note
	err := db.QueryRow(query, noteID).Scan(&note.Id, &note.Title, &note.NoteText, &note.AuthorName, &note.CreationDate, &note.IsPublic)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func AddNewNote(noteDto models.NoteDto) (*string, error) {
	query := "CALL AddNote(?, ?, ?)"

	// Переменная для хранения сгенерированного UUID заметки
	var newID string

	err := db.QueryRow(query, noteDto.Title, noteDto.Note, noteDto.Author).Scan(&newID)
	if err != nil {
		return nil, err
	}

	return &newID, nil
}

func CloseDb() {
	db.Close()
}
