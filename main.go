package main

import (
	"fmt"
	"golang-dico/dictionary"
	"sync"
	"time"
)

func main() {
	dict := dictionary.NewDictionary()
	dict.StartConcurrentOperations()

	var wg sync.WaitGroup

	// Opérations d'ajout simultanées
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			entry := dictionary.Entry{
				Key:   fmt.Sprintf("Key%d", index),
				Value: fmt.Sprintf("Value%d", index),
			}
			dict.Add(entry)
		}(i)
	}

	// Opérations de suppression simultanées
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("Key%d", index)
			dict.Remove(key)
		}(i)
	}

	// Attendez que toutes les goroutines se terminent
	wg.Wait()

	// Pause pour permettre la synchronisation
	time.Sleep(100 * time.Millisecond)

	// Affichez le dictionnaire final
	fmt.Println("Dictionnaire final:", dict.Entries())
}
