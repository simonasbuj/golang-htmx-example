package handlers

import (
	"golang-htmx-example/models"
	"html/template"
	"log"
	"net/http"
	"time"
)


func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	log.Println(r.Method, "request to", r.URL.Path)

	if (r.Method != http.MethodPost) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	if title == "" || author == "" {
		http.Error(w, "Missing title or author", http.StatusBadRequest)
		return
	}

	book := models.Book{Title: title, Author: author}

	t := template.Must(template.ParseFiles("templates/index.html"))
	t.ExecuteTemplate(w, "book-list-element", book)
}
