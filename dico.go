package main

import (
	"fmt"
	"sort"
)

// Dictionary est le dictionnaire avec des mots et des définitions.
type Dictionary map[string]string

// crée un nouveau dictionnaire vide.
func NewDictionary() Dictionary {
	return make(Dictionary)
}

// Add pour ajouter un mot et sa définition au dictionnaire.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Get récupère la définition d'un mot du dictionnaire.
func (d Dictionary) Get(word string) (string, bool) {
	definition, exists := d[word]
	return definition, exists
}

// Remove supprime un mot et sa définition du dictionnaire.
func (d Dictionary) Remove(word string) {
	delete(d, word)
}

// List renvoie une liste triée des mots et de leurs définitions.
func (d Dictionary) List() []string {
	var result []string
	for word, definition := range d {
		result = append(result, fmt.Sprintf("%s: %s", word, definition))
	}
	sort.Strings(result)
	return result
}

func main() {
	// Créez un nouveau dictionnaire.
	dictionary := NewDictionary()

	// Ajout mot et définition
	dictionary.Add("pomme", "un fruit")
	dictionary.Add("java", "un langage de programmation orienté objet")
	dictionary.Add("Putty", "Logiciel pour se connecter en SSH à un serveur distant")

	//  afficher la définition d'un mot spécifique.
	definition, exists := dictionary.Get("golang")
	if exists {
		fmt.Printf("Definition de'putty': %s\n", definition)
	} else {
		fmt.Println("Word not found in the dictionary.")
	}

	// supprimer un mot du dictionnaire.
	dictionary.Remove("apple")

	//  liste triée des mots et de leurs définitions.
	list := dictionary.List()
	fmt.Println("Dictionary entries:")
	for _, entry := range list {
		fmt.Println(entry)
	}
}
