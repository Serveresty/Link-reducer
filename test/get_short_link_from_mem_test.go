package test

import (
	"OZONTestCaseLinks/database"
	"testing"
)

func TestGetShortLinkFromMemSuccess(t *testing.T) {
	database.NewCache()
	link1, shortLink := "https://vk.com/feed", "ziT8WvYJRM"
	database.Cch.SetLink(shortLink, link1)

	shrtURL, ok := database.Cch.SearchURL(link1)
	if !ok {
		t.Fatal("error: not found link key")
	}

	if shrtURL != shortLink {
		t.Fatal("error: wrong link key")
	}
}
