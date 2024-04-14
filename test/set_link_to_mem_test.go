package test

import (
	"OZONTestCaseLinks/database"
	"testing"
)

func TestSetLinkToMemSuccess(t *testing.T) {
	database.NewCache()
	link, shortLink := "https://vk.com/feed", "ziT8WvYJRM"

	database.Cch.SetLink(shortLink, link)
	links, ok := database.Cch.Links[shortLink]
	if !ok {
		t.Fatal("error: didn't set link to cache")
	}

	if links.URL != link {
		t.Fatal("error: wrong link value")
	}
}

func TestSetLinkToMemWrongValue(t *testing.T) {
	database.NewCache()
	link1, link2, shortLink := "https://vk.com/feed", "https://vk.com/im", "ziT8WvYJRM"

	database.Cch.SetLink(shortLink, link1)
	database.Cch.SetLink(shortLink, link2)
	link, ok := database.Cch.Links[shortLink]
	if !ok {
		t.Fatal("error: didn't set link to cache")
	}

	if link.URL == link1 {
		t.Fatal("error: wrong link value")
	}
}

func TestSetLinkToMemWrongValueSuccess(t *testing.T) {
	database.NewCache()
	link1, link2, shortLink := "https://vk.com/feed", "https://vk.com/im", "ziT8WvYJRM"

	_, ok := database.Cch.GetLink(shortLink)
	if !ok {
		database.Cch.SetLink(shortLink, link1)
	}

	_, ok = database.Cch.GetLink(shortLink)
	if !ok {
		database.Cch.SetLink(shortLink, link2)
	}

	link, ok := database.Cch.Links[shortLink]
	if !ok {
		t.Fatal("error: didn't set link to cache")
	}

	if link.URL != link1 {
		t.Fatal("error: wrong link value")
	}
}
