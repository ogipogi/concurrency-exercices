Ziel: Erstelle ein Programm, in dem mehrere Schreib-Goroutinen berühmte Zitate an eine Lesegoroutine über einen Kanal senden. Der Leser gibt jedes Zitat zusammen mit der ID des Schreibers aus.

Hauptthemen: Goroutinen, Channels

Quotes:
```
quotes := [][]string{
		{"To be or not to be.", "The pen is mightier than the sword."},
		{"All that glitters is not gold.", "A rose by any other name."},
	}
```

Expected Output:
```
Writer 1: All that glitters is not gold.
Writer 1: A rose by any other name.
Writer 0: To be or not to be.
Writer 0: The pen is mightier than the sword.
```
