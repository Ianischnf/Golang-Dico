package main

import (
	"fmt"
	"./dictionary"
)

func main() {
	// Créez un nouveau dictionnaire.
	dict := dictionary.NewDictionary()

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
