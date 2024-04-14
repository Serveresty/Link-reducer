package test

import (
	"OZONTestCaseLinks/database"
	"testing"
)

func TestGetLinkFromMemSuccess(t *testing.T) {
	database.NewCache()
	link1, shortLink := "https://vk.com/feed", "ziT8WvYJRM"
	database.Cch.SetLink(shortLink, link1)
	res, ok := database.Cch.GetLink(shortLink)
	if !ok {
		t.Fatal("error: link not found")
	}

	if res != link1 {
		t.Fatal("error: wrong link value")
	}
}
