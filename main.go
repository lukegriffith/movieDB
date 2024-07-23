package main

import (
	"fmt"
	"github.com/lukegriffith/movieDB/pkg/movies"
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/index", handleDynamicTemplates)
	http.HandleFunc("/clicked", handleDynamicTemplates)
	http.HandleFunc("/movie", handleMovies)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, templatePath string, data interface{}) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

func handleDynamicTemplates(w http.ResponseWriter, r *http.Request) {
	// Using the path to determine which template to render.
	templatePath := "./templates/" + r.URL.Path + ".html"

	// For demonstration purposes, the data is kept constant. You can make this dynamic too.
	data := map[string]string{
		"Message": "This is a dynamic template!",
	}
	err := renderTemplate(w, templatePath, data)
	if err != nil {
		// If the template is not found, it's a 404. Otherwise, it's a 500.
		if _, ok := err.(*template.Error); ok {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}

func handleMovies(w http.ResponseWriter, r *http.Request) {
	// Using the path to determine which template to render.
	templatePath := "./templates/movies.html"

	switch r.Method {
	case http.MethodGet:
		// For demonstration purposes, the data is kept constant. You can make this dynamic too.
		data := map[string][]movies.Movie{
			"Movies": *movies.GetMovies(),
		}
		err := renderTemplate(w, templatePath, data)
		if err != nil {
			// If the template is not found, it's a 404. Otherwise, it's a 500.
			if _, ok := err.(*template.Error); ok {
				http.NotFound(w, r)
			} else {
				http.Error(w, "Failed to render template", http.StatusInternalServerError)
			}
		}

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		movieTitle := r.FormValue("title")
		err = movies.AddMovie(movies.MovieRequestProps{Title: movieTitle})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Movie Added")

	}
}
