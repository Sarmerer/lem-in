package visualizer

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("../visualizer/*.html"))
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
func StartServer() {
	router := http.NewServeMux()
	port := ":4243"

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../visualizer/static"))))
	router.HandleFunc("/", indexHandler)

	log.Println("Starting server, go to http://localhost" + port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
