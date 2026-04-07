package main

import (
	"log"
	"serverGo/internal/app"
	"serverGo/internal/db"
)

func main() {
	_, conn, cleanup, err := app.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	if err := db.RunSeed(conn); err != nil {
		log.Fatal(err)
	}

	log.Println("seed completed")
}
