package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	DSN   string
	DB    *sql.DB
	Store Store
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := application{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=products sslmode=disable timezone=UTC connect_timeout=5", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
	}

	conn, err := app.connectDB()
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	app.DB = conn
	app.Store = NewPostgresStore(conn)

	log.Println("Starting server on :8080...")

	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal("err con routes:", err)
	}

}
