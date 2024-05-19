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

	server.Run()
	if err := server.Run(); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}

	storage.CloseDb()
}
