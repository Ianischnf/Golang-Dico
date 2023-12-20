// dictionary.go
package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

// Structure Dictionary
type Dictionary struct {
	entries    map[string]string
	mu         sync.Mutex
	addChan    chan Entry  // Canal pour l'opération d'ajout
	removeChan chan string // Canal pour l'opération de suppression
}

// Structure Entry
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NewDictionary crée un nouveau dictionnaire à partir d'un fichier
func NewDictionary(filename string) (*Dictionary, error) {
	d := &Dictionary{
		entries:    make(map[string]string),
		addChan:    make(chan Entry),
		removeChan: make(chan string),
	}

	// Charge les entrées à partir du fichier, s'il existe
	err := d.loadFromFile(filename)
	if err != nil {
		return nil, err
	}

	// Lance les goroutines pour les opérations concurrentes
	d.StartConcurrentOperations()

	return d, nil
}

// Ajoute une entrée au dictionnaire
func (d *Dictionary) Add(entry Entry) {
	d.addChan <- entry
}

// Remove supprime une entrée du dictionnaire
func (d *Dictionary) Remove(key string) {
	d.removeChan <- key
}

// Entries renvoie une copie des entrées du dictionnaire
func (d *Dictionary) Entries() map[string]string {
	d.mu.Lock()
	defer d.mu.Unlock()

	result := make(map[string]string, len(d.entries))
	for k, v := range d.entries {
		result[k] = v
	}

	return result
}

// StartConcurrentOperations lance des goroutines pour les opérations concurrentes d'ajout et de suppression
func (d *Dictionary) StartConcurrentOperations() {
	go func() {
		for {
			select {
			case entry := <-d.addChan:
				d.mu.Lock()
				// Vérifie si la clé existe déjà dans le dictionnaire
				if _, exists := d.entries[entry.Key]; exists {
					// La clé existe, met à jour la valeur associée
					d.entries[entry.Key] = entry.Value
				} else {
					// La clé n'existe pas, ajoute une nouvelle entrée
					d.entries[entry.Key] = entry.Value
				}
				d.mu.Unlock()
				d.saveToFileAsync("dictionary.json")
			case key := <-d.removeChan:
				d.mu.Lock()
				delete(d.entries, key)
				d.mu.Unlock()
				d.saveToFileAsync("dictionary.json")
			}
		}
	}()
}

// Sauvegarde le dictionnaire dans un fichier de manière asynchrone
func (d *Dictionary) saveToFileAsync(filename string) {
	go func() {
		d.mu.Lock()
		defer d.mu.Unlock()

		data, err := json.Marshal(d.entries)
		if err != nil {
			fmt.Println("Erreur lors de la sérialisation JSON :", err)
			return
		}

		err = ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		}
	}()
}

// Charge le dictionnaire à partir d'un fichier
func (d *Dictionary) loadFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// Le fichier n'existe pas, cela peut être ignoré.
		return nil
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	err = json.Unmarshal(data, &d.entries)
	if err != nil {
		return err
	}

	return nil
}

// Get renvoie la définition associée à une clé donnée
func (d *Dictionary) Get(key string) (string, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	value, exists := d.entries[key]
	return value, exists
}

// Liste triée des mots et de leurs définitions
func (d *Dictionary) List() []Entry {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Utiliser une map temporaire pour éviter les doublons
	tempMap := make(map[string]string)

	for key, value := range d.entries {
		tempMap[key] = value
	}

	var list []Entry
	for key, value := range tempMap {
		list = append(list, Entry{Key: key, Value: value})
	}

	return list
}
