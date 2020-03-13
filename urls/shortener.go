package urls

import (
	"net/http"
	"./generator"
	"../config"
	"../html"
)

var urlID = 0

// Shortener - takes url and save shortened.
func Shortener(w http.ResponseWriter, r *http.Request) {
	originalURL := r.FormValue("url")

	var url URL

	url.SetOriginal(originalURL)
	shortenedURL := CreateURL()
	url.SetShortened(shortenedURL)
	url.SetID(urlID)
	urlID++

	url.Save()

	html.DrawIndexPage(w, r, url.ShortenedURL)
}

// CreateURL - return shortened url.
func CreateURL() string {
	generatedPath := generator.Path()

	resURL := "http://" + config.HostName + config.Port + "/" + generatedPath
	return resURL
}
