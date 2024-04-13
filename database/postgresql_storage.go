package database

import (
	"OZONTestCaseLinks/configs"
	"context"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func DbInit(cfg configs.Config) error {
	dbUrl := "postgres://" + cfg.DbUsername + ":" + cfg.DbPassword + "@" + cfg.DbHost + ":" + cfg.DbPort + "/" + cfg.DbName
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return err
	}
	DB = conn

	return nil
}

func CreateBaseTables() error {
	_, err := DB.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS "links" (id serial PRIMARY KEY, original_link VARCHAR(255) UNIQUE NOT NULL, short_link VARCHAR(10) UNIQUE NOT NULL);`)
	if err != nil {
		return err
	}
	return nil
}

func GetOriginalLinkFromDB(link string) (string, error) {
	row := DB.QueryRow(context.Background(), `SELECT original_link FROM links WHERE short_link = $1`, link)

	var resultLink string
	err := row.Scan(&resultLink)
	if err != nil {
		return "", err
	}

	return resultLink, nil
}

func GetReducedLinkFromDB(link string) (string, error) {
	row := DB.QueryRow(context.Background(), `SELECT short_link FROM links WHERE original_link = $1`, link)

	var resultLink string
	err := row.Scan(&resultLink)
	if err != nil {
		return "", err
	}

	return resultLink, nil
}

func SetReducedLinkToDB(shortLink, link string) error {
	_, err := DB.Exec(context.Background(), `INSERT INTO "links" (original_link, short_link) VALUES($1, $2)`, link, shortLink)
	if err != nil {
		return err
	}
	return nil
}
