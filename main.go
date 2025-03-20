package main

import (
	/*
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
	"fmt"
	*/
	"html/template"
	"log"
	"net/http"
)
/*
var db *sql.DB

type HomeData struct {
	Connecte bool
	Message  string
}
*/
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
	/*
	data := HomeData{
		Connecte: false,
	}

	if r.Method == http.MethodPost {
		data.Connecte = true
	}
	renderTemplate(w, "site", data)
	*/
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
/*
func inscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Récupérer les données du formulaire
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Vérifier que les champs ne sont pas vides
		if username == "" || password == "" {
			http.Error(w, "Veuillez remplir tous les champs", http.StatusBadRequest)
			return
		}

		// Insérer les données dans la base SQLite
		_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
		if err != nil {
			http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
			return
		}

		// Rediriger vers la page de connexion
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}

	// Si ce n'est pas une requête POST, afficher la page d'inscription
	renderTemplate(w, "inscription", nil)
}

func connexionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Récupérer les données du formulaire
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Vérifier que les champs ne sont pas vides
		if username == "" || password == "" {
			http.Error(w, "Veuillez remplir tous les champs", http.StatusBadRequest)
			return
		}

		// Vérifier si l'utilisateur existe dans la base de données
		var storedPassword string
		err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
		if err != nil {
			http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
			return
		}

		// Vérifier le mot de passe (en vrai, il faudrait le hacher avec bcrypt)
		if password != storedPassword {
			http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
			return
		}

		// Connexion réussie
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Si ce n'est pas une requête POST, afficher la page de connexion
	renderTemplate(w, "connexion", nil)
}

func initDB() {
	var err error
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	// Création de la table users si elle n'existe pas
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        password TEXT
    )`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base de données initialisée avec succès")
}
*/
func main() {
	/*
	initDB()
	defer db.Close()
	*/

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	
	http.HandleFunc("/", homeHandler)

	port := ":8080"
	log.Println("Serveur démarré sur http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}