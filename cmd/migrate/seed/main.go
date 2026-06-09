package main

import (
	"log"
	"os"

	"github.com/Leonfarhan/simple-social/internal/db"
	"github.com/Leonfarhan/simple-social/internal/store"
)

func main() {
	conn, err := db.New(os.Getenv("DB_ADDR"), "15m", 3, 3)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	storage := store.NewPostgresStorage(conn)

	db.Seed(storage, conn)
}
