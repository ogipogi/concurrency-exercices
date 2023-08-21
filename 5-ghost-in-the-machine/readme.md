Übung 5: Geist in der Maschine

Ziel: Simuliere einen schelmischen Computergeist, der zufällig ein gemeinsames Dokument (eine Zeichenfolge) verändert. Mehrere Schreib-Goroutinen repräsentieren Personen, die an dem Dokument arbeiten, und eine Geist-Goroutine fügt zufällig lustige Wörter oder Phrasen hinzu. Verwende Kanäle zur Kommunikation von Änderungen und einen Mutex, um das gemeinsame Dokument zu schützen.

Hauptthemen: Goroutinen, Kanäle, Mutex

Expected Output:
```
This is an important document. Writer 1 added. Boo! 
Writer 0 added. Boo! Writer 0 added. Boo! Writer 0 
added. Boo! Writer 0 added. Boo! Writer 1 added. 
Boo! Writer 0 added. Boo! Writer 0 added. Boo! 
Writer 1 added. Boo! Writer 0 added. Boo! 
Writer 1 added. Boo! Writer 0 added. Boo! 
Writer 1 added. Boo! Writer 1 added. Boo! 
Writer 0 added. Boo! Writer 1 added. Boo! 
Writer 0 added. Boo! Writer 1 added. Boo! 
Writer 0 added. Boo! Writer 0 added. Boo! 
Writer 0 added. Boo! Writer 1 added. Boo! 
Writer 1 added. Boo! Writer 0 added. Boo! 
Writer 0 added. Boo! 

```