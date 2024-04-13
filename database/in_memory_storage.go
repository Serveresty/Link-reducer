package database

import (
	"sync"
	"time"
)

var Cch *Cache

type Cache struct {
	mtx   sync.Mutex
	Links map[string]Link
}

type Link struct {
	URL     string
	Created time.Time
}

func NewCache() {
	links := make(map[string]Link)
	cache := &Cache{
		Links: links,
	}

	Cch = cache
}

func (c *Cache) SetLink(shortLink string, link string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.Links[shortLink] = Link{
		URL:     link,
		Created: time.Now(),
	}
}

func (c *Cache) GetLink(shortLink string) (string, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	link, ok := c.Links[shortLink]
	if !ok {
		return "", false
	}

	return link.URL, true
}

func (c *Cache) SearchURL(link string) (string, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	for short, url := range c.Links {
		if url.URL == link {
			return short, true
		}
	}
	return "", false
}
