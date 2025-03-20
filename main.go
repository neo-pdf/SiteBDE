package main

import (
	"html/template"
	"log"
	"net/http"
)

type HomeData struct {
	Connecte bool
}

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	tmpl = "templates/" + tmpl + ".html"
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		Connecte: false,
	}

	if r.Method == http.MethodPost {
		data.Connecte = true
	}

	renderTemplate(w, "site", data)
}

func postulerHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "postuler", nil)
}

func supportHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "support", nil)
}

func membresHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "membres", nil)
}

func connexionHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		Connecte: false,
	}
	if r.Method == http.MethodPost {
		data.Connecte = true
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	renderTemplate(w, "connexion", data)
}

func inscriptionHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		Connecte: false,
	}

	if r.Method == http.MethodPost {
		data.Connecte = true
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}

	renderTemplate(w, "inscription", data)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	http.HandleFunc("/", homeHandler)

	port := ":8080"
	log.Println("Serveur démarré sur http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
