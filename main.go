// main.go
package main

import (
	"encoding/json"
	"fmt"
	"golang-dico/dictionary"
	"io/ioutil"
	"net/http"
)

var dict *dictionary.Dictionary

func main() {
	// Créez un nouveau dictionnaire.
	dictionary, err := dictionary.NewDictionary("dictionary.json")
	if err != nil {
		fmt.Println("Erreur lors de la création du dictionnaire :", err)
		return
	}
	dict = dictionary

	//la fonction pour créer le fichier JSON
	createJSONFile()
	// Routes
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/remove", removeHandler)
	http.HandleFunc("/list", listHandler)

	// Démarrer le serveur HTTP
	err = http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur HTTP :", err)
	}
}

// addHandler gère les requêtes POST pour ajouter une entrée au dictionnaire.
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Décodage du corps de la requête (JSON) pour obtenir une entrée.
	var entry dictionary.Entry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Ajouter l'entrée au dictionnaire.
	dict.Add(entry)
	w.WriteHeader(http.StatusCreated)

	// Appeler la sauvegarde après l'ajout
	dict.SaveToFile("dictionary.json")

	// Appeler la fonction pour créer le fichier JSON après avoir ajouté une entrée
	createJSONFile()

}

// getHandler gère les requêtes GET pour récupérer une définition par mot.
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Obtenir le mot à partir des paramètres de requête.
	word := r.URL.Query().Get("word")
	if word == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Récupérer la définition du mot dans le dictionnaire.
	definition, exists := dict.Get(word)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Répondre avec la définition.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(definition))
}

// removeHandler gère les requêtes DELETE pour supprimer une entrée par mot.
func removeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Obtenir le mot à partir des paramètres de requête.
	word := r.URL.Query().Get("word")
	if word == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Supprimer l'entrée du dictionnaire.
	dict.Remove(word)
	fmt.Printf("Suppression réussie de la clé : %s\n", word)

	w.WriteHeader(http.StatusOK)
}

// listHandler gère les requêtes GET pour afficher toutes les entrées du dictionnaire.
func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Obtenir la liste triée des entrées du dictionnaire.
	list := dict.List()

	// Convertir la liste en format JSON.
	response, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Répondre avec la liste JSON.
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// createJSONFile crée un fichier JSON avec les mots du dictionnaire
func createJSONFile() {
	// Obtenir la liste triée des entrées du dictionnaire.
	list := dict.List()

	// Convertir la liste en format JSON.
	data, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation JSON :", err)
		return
	}

	// Écrire les données JSON dans un fichier.
	err = ioutil.WriteFile("dictionary_export.json", data, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier JSON :", err)
	}
}
