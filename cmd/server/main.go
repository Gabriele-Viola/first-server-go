package main

import (
	"log"
	"net/http"
	"serverGo/internal/routes"

	"serverGo/internal/app"
)

func main() {

	_, conn, cleanup, err := app.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	mux := routes.Configure(conn)

	log.Println("Server avviato su :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
