package main

import (
	"OZONTestCaseLinks/configs"
	"OZONTestCaseLinks/database"
	"OZONTestCaseLinks/internal/transport"
	"flag"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	storage := flag.String("storage", "postgresql", "Тип хранилища postgresql/in-memory")
	flag.Parse()

	switch *storage {
	case "postgresql":
		postgresHandler()
	case "in-memory":
		inMemHandler()
	default:
		log.Fatal("Wrong storage param")
	}

	database.Storage = *storage
	mux := http.NewServeMux()
	transport.Routs(mux)

	addr := configs.LoadServerConfig()

	if err := http.ListenAndServe(addr.Host+addr.Port, mux); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}

func postgresHandler() {
	cfg := configs.LoadConfig()
	err := database.DbInit(cfg)
	if err != nil {
		log.Fatalf("Error connecting DB: %s", err.Error())
	}

	err = database.CreateBaseTables()
	if err != nil {
		log.Fatalf("Error creating base tables: %s", err.Error())
	}
}

func inMemHandler() {
	database.NewCache()
}
