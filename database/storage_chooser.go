package database

import (
	"OZONTestCaseLinks/pkg"
	"fmt"
)

var Storage string

func ReduceLink(link string) (string, error) {
	if Storage == "postgresql" {
		oLink, err := GetReducedLinkFromDB(link)
		if err != nil {
			for {
				shortLink := pkg.GenerateLink(10)
				_, err := GetOriginalLinkFromDB(link)
				if err != nil {
					err := SetReducedLinkToDB(shortLink, link)
					if err != nil {
						return "", err
					}
					return shortLink, nil
				}
			}
		}
		return oLink, nil
	}

	rLink, ok := Cch.SearchURL(link)
	if ok {
		return rLink, nil
	}

	for {
		shortLink := pkg.GenerateLink(10)
		_, ok := Cch.GetLink(shortLink)
		if !ok {
			Cch.SetLink(shortLink, link)
			return shortLink, nil
		}
	}
}

func OriginalLink(link string) (string, error) {
	if Storage == "postgresql" {
		originalLink, err := GetOriginalLinkFromDB(link)
		if err != nil {
			return "", err
		}
		return originalLink, nil
	}

	originalLink, ok := Cch.GetLink(link)
	if !ok {
		return "", fmt.Errorf("link not found")
	}
	return originalLink, nil
}
