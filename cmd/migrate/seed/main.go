package main

import (
	"log"
	"os"

	"github.com/Leonfarhan/simple-social/internal/db"
	"github.com/Leonfarhan/simple-social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := db.New(os.Getenv("DB_ADDR"), "3s", 3, 15)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewPostgresStorage(conn)
	if err := db.Seed(store); err != nil {
		log.Fatal(err)
	}
}
