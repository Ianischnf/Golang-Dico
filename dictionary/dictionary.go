// dictionary.go
package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

// Dictionary est le dictionnaire avec des mots et des définitions.
type Dictionary map[string]string

// crée un nouveau dictionnaire vide.
func NewDictionary(filename string) (Dictionary, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier : %v", err)
	}

	var dict Dictionary
	err = json.Unmarshal(data, &dict)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la conversion du JSON en dictionnaire : %v", err)
	}

	return dict, nil
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

// SaveToFile enregistre le dictionnaire dans un fichier JSON.
func (d Dictionary) SaveToFile(filename string) error {
	// Convertir le dictionnaire en JSON.
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion du dictionnaire en JSON : %v", err)
	}

	// Écrire le JSON dans le fichier.
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier : %v", err)
	}

	return nil
}
