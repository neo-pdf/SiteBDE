package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmpl = "templates/" + tmpl + ".html"
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "site")
}

func postulerHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "postuler")
}

func supportHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "support")
}

func membresHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "membres")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/postuler", postulerHandler)
	http.HandleFunc("/support", supportHandler)
	http.HandleFunc("/membres", membresHandler)

	port := ":8080"
	log.Println("Serveur démarré sur http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
