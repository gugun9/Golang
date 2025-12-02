package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/about.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/contact.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	})

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}