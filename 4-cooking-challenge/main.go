package main

import (
	"fmt"
	"sync"
)

var pantry = map[string]int{
	"sugar":     5,
	"flour":     3,
	"chocolate": 2,
}

var pantryMutex sync.Mutex

func chef(id int, ingredientRequests chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Reduziert den Zähler des WaitGroups, wenn die Funktion zurückkehrt
	ingredients := []string{"sugar", "flour", "chocolate"}
	for _, ingredient := range ingredients {
		ingredientRequests <- ingredient
	}

	fmt.Printf("Chef %d finished the recipe!\n", id)
}

func pantryManager(ingredientRequests chan string) {
	for ingredient := range ingredientRequests {
		pantryMutex.Lock()
		if pantry[ingredient] > 0 {
			pantry[ingredient]--
			fmt.Printf("Served %s. Remaining: %d\n", ingredient, pantry[ingredient])
		}
		pantryMutex.Unlock()
	}
}

func main() {
	ingredientRequests := make(chan string)
	var wg sync.WaitGroup

	go pantryManager(ingredientRequests)

	wg.Add(2) // Fügen Sie 2 zum Zähler des WaitGroups hinzu, da es 2 Köche gibt
	for i := 0; i < 2; i++ {
		go chef(i, ingredientRequests, &wg)
	}

	wg.Wait() // Warten Sie, bis alle Köche fertig sind
	close(ingredientRequests)
}
