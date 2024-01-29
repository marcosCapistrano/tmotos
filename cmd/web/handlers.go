package main

import (
	"net/http"
	"strings"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from snippetbox"))
}

func (app *application) store(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base_layout.html",
		"./ui/html/partials/navbar.html",
		"./ui/html/pages/store.html",
	}

	fullPath := strings.Split(r.URL.Path, "/")
	_ = app.repository.SearchProducts(strings.Join(fullPath[2:], "/"))

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Debug(err.Error())
		http.Error(w, "Internal server error", 500)
	}

	err = ts.ExecuteTemplate(w, "base_layout", nil)
	if err != nil {
		app.logger.Debug(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}
