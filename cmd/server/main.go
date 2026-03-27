package main

import (
	"log"
	"net/http"
	"serverGo/internal/routes"
)

func main() {
	mux := routes.Configure()

	log.Println("Server avviato su :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
