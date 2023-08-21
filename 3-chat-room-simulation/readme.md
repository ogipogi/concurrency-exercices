Übung 3: Chatraum-Simulation

Ziel: Simuliere einen einfachen Chatraum mit mehreren Clients als Goroutinen. Jeder Client sendet zufällige Nachrichten an einen gemeinsamen Chatkanal, und eine Servergoroutine liest die Nachrichten und gibt sie mit dem Namen des Clients aus.

Hauptthemen: Goroutinen, Chanels, Mutex

Expected Output:
```
Client 1: Hello!
Client 0: Hello!
Client 2: Hello!
Client 2: Hello!
Client 1: Hello!
Client 1: Hello!
Client 0: Hello!
Client 2: Hello!
Client 1: Hello!
Client 2: Hello!
Client 2: Hello!
Client 0: Hello!
Client 1: Hello!
Client 2: Hello!
Client 1: Hello!
Client 0: Hello!
Client 2: Hello!

```