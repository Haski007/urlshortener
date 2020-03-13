package urls

import (
	"log"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"net/http"

	"../config"
	"../database"
)

type m bson.M

// MakeRedirection ...
func MakeRedirection(w http.ResponseWriter, r *http.Request) error {
	var url URL

	fullURL := "http://" + config.HostName + config.Port + r.URL.Path

	i, ok := Cache.Get(fullURL)
	if ok {
		url = i.(URL)
		http.Redirect(w, r, url.OriginalURL, 302)
		log.Printf("Redirected to [%s] found in cache, with status 302\n", url.OriginalURL)
		return nil
	}

	err := database.UrlsCollection.Find(m{"shortened" : fullURL}).One(&url)
	if err != nil {
		return fmt.Errorf("Such url does not exist")
	}

	http.Redirect(w, r, url.OriginalURL, 302)
	log.Printf("Redirected to [%s] found in database, with status 302\n", url.OriginalURL)
	return nil
}