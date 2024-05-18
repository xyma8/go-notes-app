package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	notes "github.com/xyma8/go-notes-app"
	storage "github.com/xyma8/go-notes-app/pkg/storage"
)

func main() {
	server := notes.NewAPIServer(":8000")

	dsn := "root:admin@tcp(127.0.0.1:3306)/gonotesdb"

	storage.ConnectDB(dsn)

	/*
		// Пример выполнения запроса
		var version string
		err = db.QueryRow("SELECT VERSION()").Scan(&version)
		if err != nil {
			log.Fatalf("Ошибка выполнения запроса: %v", err)
		}
		log.Printf("Версия MySQL: %s\n", version)
	*/

	server.Run()
	if err := server.Run(); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
