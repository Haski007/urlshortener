package html

import (
	"html/template"
	"net/http"

	"../errno"
)

// DrawIndexPage - draws index.html.
func DrawIndexPage(w http.ResponseWriter, r *http.Request, data string) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		errno.PrintError(err)
		return
	}

	if data == "" {
		err = t.ExecuteTemplate(w, "index", nil)
	} else {
		err = t.ExecuteTemplate(w, "index", data)
	}
	if err != nil {
		errno.PrintError(err)
		return
	}
}