package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
		return nil, err
	}

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка проверки соединения: %v", err)
		return nil, err
	}

	log.Printf("Успешное подключение к базе данных!")
	return db, nil
}
