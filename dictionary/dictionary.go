package dictionary

import (
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
	Key   string
	Value string
}

// NewDictionary crée un nouveau dictionnaire
func NewDictionary() *Dictionary {
	return &Dictionary{
		entries:    make(map[string]string),
		addChan:    make(chan Entry),
		removeChan: make(chan string),
	}
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
				d.entries[entry.Key] = entry.Value
				d.mu.Unlock()
			case key := <-d.removeChan:
				d.mu.Lock()
				delete(d.entries, key)
				d.mu.Unlock()
			}
		}
	}()
}
