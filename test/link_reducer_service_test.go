package test

import (
	"OZONTestCaseLinks/configs"
	"OZONTestCaseLinks/database"
	"OZONTestCaseLinks/internal/services"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestLinkReducerDBCreateShortLink(t *testing.T) {
	if err := godotenv.Load("../configs/.env"); err != nil {
		log.Print("No .env file found")
	}

	database.Storage = "postgresql"
	dbCfg := configs.LoadConfig()
	dbCfg.DbName = os.Getenv("DB_TEST_NAME")

	err := database.DbInit(dbCfg)
	if err != nil {
		t.Fatalf("error db init: %s", err.Error())
	}

	err = database.CreateBaseTables()
	if err != nil {
		t.Fatalf("error creating base tables: %s", err.Error())
	}

	trs, err := database.DB.Begin(context.Background())
	if err != nil {
		t.Fatalf("error start transaction: %s", err.Error())
	}

	link := "https://vk.com/feed"

	jsonResp, err := json.Marshal(link)
	if err != nil {
		t.Fatalf("error creating json payload: %s", err.Error())
	}

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonResp))
	if err != nil {
		t.Fatalf("error creating new request: %s", err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.LinkReducer)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v, returned data: %v", status, http.StatusOK, rr.Body)
	}

	if len(rr.Body.String()) < 10 {
		t.Fatalf("handler returned unexpected body: got %v", rr.Body.String())
	}

	err = trs.Rollback(context.Background())
	if err != nil {
		t.Fatalf("error rollback transaction: %s", err.Error())
	}
}

func TestLinkReducerDBGetLink(t *testing.T) {
	if err := godotenv.Load("../configs/.env"); err != nil {
		log.Print("No .env file found")
	}

	database.Storage = "postgresql"
	dbCfg := configs.LoadConfig()
	dbCfg.DbName = os.Getenv("DB_TEST_NAME")

	err := database.DbInit(dbCfg)
	if err != nil {
		t.Fatalf("error db init: %s", err.Error())
	}

	err = database.CreateBaseTables()
	if err != nil {
		t.Fatalf("error creating base tables: %s", err.Error())
	}

	trs, err := database.DB.Begin(context.Background())
	if err != nil {
		t.Fatalf("error start transaction: %s", err.Error())
	}

	link := "https://vk.com/feed"

	jsonResp, err := json.Marshal(link)
	if err != nil {
		t.Fatalf("error creating json payload: %s", err.Error())
	}

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonResp))
	if err != nil {
		t.Fatalf("error creating new request: %s", err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.LinkReducer)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v, returned data: %v", status, http.StatusOK, rr.Body)
	}

	if len(rr.Body.String()) < 10 {
		t.Fatalf("handler returned unexpected body: got %v", rr.Body.String())
	}

	path := rr.Body.String()
	path = path[2 : len(path)-1]

	req1, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatalf("error creating new request: %s", err.Error())
	}

	rr1 := httptest.NewRecorder()
	handler.ServeHTTP(rr1, req1)

	if status := rr1.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v, returned data: %v", status, http.StatusOK, rr1.Body)
	}

	result := rr1.Body.String()
	result = result[1 : len(result)-1]
	if result != link {
		t.Fatalf("handler returned wrong responce data: got %v want %s", rr1.Body, link)
	}

	err = trs.Rollback(context.Background())
	if err != nil {
		t.Fatalf("error rollback transaction: %s", err.Error())
	}
}

func TestLinkReducerMemCreateShortLink(t *testing.T) {
	database.NewCache()
	database.Storage = "in-memory"

	link := "https://vk.com/feed"

	jsonResp, err := json.Marshal(link)
	if err != nil {
		t.Fatalf("error creating json payload: %s", err.Error())
	}

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonResp))
	if err != nil {
		t.Fatalf("error creating new request: %s", err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.LinkReducer)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v, returned data: %v", status, http.StatusOK, rr.Body)
	}

	if len(rr.Body.String()) < 10 {
		t.Fatalf("handler returned unexpected body: got %v", rr.Body.String())
	}

	t.Logf("Returns result: %s", rr.Body.String())
}

func TestLinkReducerMemGetLink(t *testing.T) {
	database.NewCache()
	database.Storage = "in-memory"

	link := "https://vk.com/feed"

	jsonResp, err := json.Marshal(link)
	if err != nil {
		t.Fatalf("error creating json payload: %s", err.Error())
	}

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonResp))
	if err != nil {
		t.Fatalf("error creating new request: %s", err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.LinkReducer)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v, returned data: %v", status, http.StatusOK, rr.Body)
	}

	if len(rr.Body.String()) < 10 {
		t.Fatalf("handler returned unexpected body: got %v", rr.Body.String())
	}

	path := rr.Body.String()
	path = path[2 : len(path)-1]

	req1, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatalf("error creating new request: %s", err.Error())
	}

	rr1 := httptest.NewRecorder()
	handler.ServeHTTP(rr1, req1)

	if status := rr1.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v, returned data: %v", status, http.StatusOK, rr1.Body)
	}

	result := rr1.Body.String()
	result = result[1 : len(result)-1]
	if result != link {
		t.Fatalf("handler returned wrong responce data: got %v want %s", rr1.Body, link)
	}

	t.Logf("Returns short link result: %s", rr.Body.String())
	t.Logf("Returns link result: %s", rr1.Body.String())
}
