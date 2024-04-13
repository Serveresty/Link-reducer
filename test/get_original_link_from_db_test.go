package test

import (
	"OZONTestCaseLinks/configs"
	"OZONTestCaseLinks/database"
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetOriginalLinkFromDbSuccess(t *testing.T) {
	if err := godotenv.Load("../configs/.env"); err != nil {
		log.Print("No .env file found")
	}
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

	link, shortLink := "https://vk.com/feed", "ziT8WvYJRM"
	err = database.SetReducedLinkToDB(shortLink, link)
	if err != nil {
		err = trs.Rollback(context.Background())
		if err != nil {
			t.Fatalf("error rollback transaction: %s", err.Error())
		}
		t.Fatalf("error insert into db: %s", err.Error())
	}

	oLink, err := database.GetOriginalLinkFromDB(shortLink)
	if err != nil {
		err = trs.Rollback(context.Background())
		if err != nil {
			t.Fatalf("error rollback transaction: %s", err.Error())
		}
		t.Fatalf("error insert into db: %s", err.Error())
	}

	if oLink != link {
		err = trs.Rollback(context.Background())
		if err != nil {
			t.Fatalf("error rollback transaction: %s", err.Error())
		}
		t.Fatal("error wrong original link result")
	}

	err = trs.Rollback(context.Background())
	if err != nil {
		t.Fatalf("error rollback transaction: %s", err.Error())
	}
}
