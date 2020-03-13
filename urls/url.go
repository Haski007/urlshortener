package urls

import (
	"time"
	"fmt"
	"../database"
	"github.com/patrickmn/go-cache"

	"../config"
)

// URL - the main entity struct.
type URL struct {
	ID				int	`json:"ID" bson:"_id"`
	OriginalURL		string `json:"OriginalURL" bson:"original"`
	ShortenedURL	string `json:"ShortenedURL" bson:"shortened"`
	Created			time.Time `json:"-" bson:"-"`
}

// Cache - application cache with timeout in 2 hours
var Cache = cache.New(config.Expiration, config.Expiration)

// Save - pushes url struct to cache and mongoDB.
func (url URL) Save() {
	database.UrlsCollection.Insert(url)

	Cache.Set(url.ShortenedURL, url, cache.DefaultExpiration)
	fmt.Printf("Url: [%s] and [%s], has beed saved at memory and mongoDB\n", url.ShortenedURL, url.OriginalURL)
}

// SetCurrentTime set current time.
func (url *URL) SetCurrentTime() {
	url.Created = time.Now()
}

// GetOriginal - returns original url.
func (url URL) GetOriginal() string {
	return url.OriginalURL
}

// GetShortened - returns shortened url.
func (url URL) GetShortened() string {
	return url.OriginalURL
}

// SetOriginal - set url.OriginalUrl by your data.
func (url *URL) SetOriginal(str string) {
	url.OriginalURL = str
}

// SetShortened - set url.ShortenedUrl by your data
func (url *URL) SetShortened(str string) {
	url.ShortenedURL = str
}

// SetID - set id of url.
func (url *URL) SetID(id int) {
	url.ID = id
}