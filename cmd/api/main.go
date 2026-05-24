package main

import (
	"log"
	"os"

	"github.com/Leonfarhan/simple-social/internal/db"
	"github.com/Leonfarhan/simple-social/internal/store"
	"github.com/joho/godotenv"
)

const version = "0.0.1"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: os.Getenv("ADDR"),
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
		env: os.Getenv("ENV"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxIdleTime,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	app := &application{
		config: cfg,
		store:  store.NewPostgresStorage(db),
	}

	log.Fatal(app.run(app.mount()))
}
