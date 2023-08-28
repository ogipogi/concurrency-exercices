package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var document = "This is an important document. "
var documentMutex sync.Mutex

func writer(id int, updates chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		documentMutex.Lock()
		document += fmt.Sprintf("Writer %d added. ", id)
		documentMutex.Unlock()
		updates <- "update"
	}
}

func ghost(updates chan string) {
	for range updates {
		documentMutex.Lock()
		document += "Boo! "
		documentMutex.Unlock()
	}
}

func main() {
	updates := make(chan string)

	go ghost(updates)

	for i := 0; i < 2; i++ {
		go writer(i, updates)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(document)
}
