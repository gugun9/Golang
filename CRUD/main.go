package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tmpl *template.Template

// Data Model
type Item struct {
	ID    int
	Title string
}

// In-memory storage (tanpa database)
var items = []Item{
	{ID: 1, Title: "Belajar Golang"},
	{ID: 2, Title: "Belajar Tailwind"},
}

// Load template sekali saja
func init() {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		panic("ERROR Load Template: " + err.Error())
	}
}

// Home
func home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", items)
}

// Create Form
func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		id := len(items) + 1
		items = append(items, Item{ID: id, Title: title})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "create.html", nil)
}

// Edit
func edit(w http.ResponseWriter, r *http.Request) {
	// Ambil ID
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	// Temukan data
	var data Item
	for _, v := range items {
		if v.ID == id {
			data = v
			break
		}
	}

	if r.Method == "POST" {
		title := r.FormValue("title")
		for i := range items {
			if items[i].ID == id {
				items[i].Title = title
				break
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl.ExecuteTemplate(w, "edit.html", data)
}

// Delete
func delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for i := range items {
		if items[i].ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", home)
	http.HandleFunc("/create", create)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/delete", delete)

	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}