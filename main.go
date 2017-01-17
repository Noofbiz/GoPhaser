package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// this serves the index page to test the game on
	http.HandleFunc("/", mainHandler)

	// this gets the libraries for your game
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// listen and serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()

	p := filepath.Join(cwd, "index.html")

	t := template.Must(template.ParseFiles(p))

	if err := t.Execute(w, nil); err != nil {
		log.Print(err.Error())
	}
}
