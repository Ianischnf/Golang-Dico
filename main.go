// main.go
package main

import (
	"fmt"

	"golang-dico/dictionary"
)

func main() {
	// Créez un nouveau dictionnaire.
	dict, err := dictionary.NewDictionary("dictionary.json")
	if err != nil {
		fmt.Println("Erreur lors de la création du dictionnaire :", err)
		return
	}

	// Ajout mot et définition
	dict.Add("pomme", "un fruit")
	dict.Add("java", "un langage de programmation orienté objet")
	dict.Add("Putty", "Logiciel pour se connecter en SSH à un serveur distant")

	//  afficher la définition d'un mot spécifique.
	definition, exists := dict.Get("golang")
	if exists {
		fmt.Printf("Definition de 'putty': %s\n", definition)
	} else {
		fmt.Println("Word not found in the dictionary.")
	}

	// supprimer un mot du dictionnaire.
	dict.Remove("apple")

	//  liste triée des mots et de leurs définitions.
	list := dict.List()
	fmt.Println("Dictionary entries:")
	for _, entry := range list {
		fmt.Println(entry)
	}
}
