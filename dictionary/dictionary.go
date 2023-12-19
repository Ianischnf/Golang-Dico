package dictionary

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
