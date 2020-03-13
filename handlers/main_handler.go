package handlers

import (
	"io"
	"net/http"

	"../errno"
	"../urls"
	"../html"
)

// Main handler function by path "/".
func Main(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/" {
			html.DrawIndexPage(w, r, "")
			return
		} 
		err := urls.MakeRedirection(w, r)
		if err != nil {
			errno.PrintError(err)
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, `{"alive": true}`)
			return
		}
	} else {
		urls.Shortener(w, r)
	}
}