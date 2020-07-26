package main

import (
	"log"
	"net/http"
	"text/template"
)

var indexTpl *template.Template = template.Must(template.ParseGlob("templates/index/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//404
	}
	switch r.Method {
	case "GET":
		indexTpl.ExecuteTemplate(w, "index.html", nil)
	default:
		//wrong method
	}
}

func main() {
	router := http.NewServeMux()
	port := ":4243"

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../ui/static"))))
	router.HandleFunc("/", index)

	log.Println("Starting server, go to http://localhost" + port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
