package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"text/template"
)

var templates *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("*.html"))
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

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	// case "linux":
	// 	err = exec.Command("xdg-open", url).Start()
	case "windows", "linux":
		err = exec.Command("explorer.exe", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	//openbrowser("index.html")
	router := http.NewServeMux()
	port := ":4243"

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("/", indexHandler)

	log.Println("Starting server, go to http://localhost" + port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
