Übung 4: Kochwettbewerb

Ziel: Erstelle eine Kochwettbewerbssimulation, in der mehrere Chef-Goroutinen versuchen, ein Rezept zu vervollständigen. Jeder Schritt des Rezepts erfordert eine bestimmte Zutat aus der Speisekammer. Verwende einen Kanal, um Zutaten anzufragen und zu empfangen, und einen Mutex, um sicherzustellen, dass nur ein Chef gleichzeitig auf die Speisekammer zugreift.

Hauptthemen: Goroutinen, Channels, Mutex, WaitGroup

Expected Output:
```
Served sugar. Remaining: 4
Served sugar. Remaining: 3
Served flour. Remaining: 2
Chef 1 finished the recipe!
Served chocolate. Remaining: 1
Served flour. Remaining: 1
Served chocolate. Remaining: 0
Chef 0 finished the recipe!

```