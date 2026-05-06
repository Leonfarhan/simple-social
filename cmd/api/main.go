package main

import (
	"log"
	"os"

	"github.com/Leonfarhan/simple-social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: os.Getenv("ADDR"),
	}

	app := &application{
		config: cfg,
		store: store.NewPostgresStorage(nil),
	}

	log.Fatal(app.run(app.mount()))
}
