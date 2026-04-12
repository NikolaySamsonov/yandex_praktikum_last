package main

import (
	"log"
	"net/http"
	"os"

	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
)

func main() {
	port := "7540"
	if envPort := os.Getenv("TODO_PORT"); envPort != "" {
		port = envPort
	}

	dbFile := "scheduler.db"
	if envDB := os.Getenv("TODO_DBFILE"); envDB != "" {
		dbFile = envDB
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatal("DB init error:", err)
	}
	defer db.DB.Close()

	api.InitAuth(os.Getenv("TODO_PASSWORD"))
	api.Init()

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	log.Println("Server started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
