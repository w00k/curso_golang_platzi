package main

import (
	"fmt"
)

// chanel como parametros de entrada o salida, es buena practica especificar si es de entrada o salida
// para salida es con 'c chan <- string'
// para entrada es con 'c <- chan string'
func say(text string, c chan<- string) {
	c <- text
}

// channels soporta enviar datos entre distintos goroutines.
// goroutines son más eficientes que usar channels, pero no permiten comunicación entre otras goroutines
func main() {
	c := make(chan string, 1) //  1 cantidad límite

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)
}
