package handlers

import (
	"html/template"
	"log"
	"net/http"

	"golang-htmx-example/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, "request to", r.URL.Path)

	t := template.Must(template.ParseFiles("templates/index.html"))

	books := map[string][]models.Book{
		"Books": {
			{Title: "Da Vinki", Author: "Da Vinki Brothers"},
			{Title: "Stormlight Archive", Author: "Brando"},
			{Title: "KET rules", Author: "Gitan Nawsaeda"},
		},
	}

	t.Execute(w, books)
}
