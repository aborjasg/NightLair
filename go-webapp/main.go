package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "My Go Web App",
	}
	templates.ExecuteTemplate(w, "index.html", data)
}

func main() {
	// Serve static files (JS, CSS, images)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", homeHandler)

	log.Println("Web app running on :8081")
	http.ListenAndServe(":8081", nil)
}
