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

	// Créez des instances de la structure Entry avec les données spécifiées
	entry1 := dictionary.Entry{Key: "pomme", Value: "un fruit"}
	entry2 := dictionary.Entry{Key: "java", Value: "un langage de programmation orienté objet"}
	entry3 := dictionary.Entry{Key: "Putty", Value: "Logiciel pour se connecter en SSH à un serveur distant"}

	// Ajoutez ces entrées au dictionnaire en utilisant la méthode Add
	dict.Add(entry1)
	dict.Add(entry2)
	dict.Add(entry3)

	// Afficher la définition d'un mot spécifique.
	definition, exists := dict.Get("java")
	if exists {
		fmt.Printf("Definition de 'java': %s\n", definition)
	} else {
		fmt.Println("Word not found in the dictionary.")
	}

	// Supprimer un mot du dictionnaire.
	dict.Remove("java")

	// Liste triée des mots et de leurs définitions.
	list := dict.List()
	fmt.Println("Dictionary entries:")
	for _, entry := range list {
		fmt.Println(entry)
	}
}
