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
	log.Println("Service started")
	storage := flag.String("STORAGE", "postgresql", "Тип хранилища postgresql/in-memory")
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
	log.Println("listening " + addr.Host + ":" + addr.Port)
	if err := http.ListenAndServe(addr.Host+":"+addr.Port, mux); err != nil {
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
	log.Println("postgresql started")
}

func inMemHandler() {
	database.NewCache()
	log.Println("in-memory started")
}
