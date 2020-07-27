package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template = template.Must(template.ParseGlob("../templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, 404)
		return
	}
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "index.html", nil)
	default:
		errorHandler(w, r, 405)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, errorCode int) {
	w.WriteHeader(errorCode)
	templates.ExecuteTemplate(w, "error.html", errorCode)
}

func main() {
	router := http.NewServeMux()
	port := ":4243"

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	router.HandleFunc("/", indexHandler)

	log.Println("Starting server, go to http://localhost" + port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}