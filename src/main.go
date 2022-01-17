package main

import "fmt"

func main() {
	//
	helloMessage := "hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos) //no hay certeza del tipo de dato, por eso %v

	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// conocer tl tipo de dato de una variable
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("helloMessage: %T \n", cursos)
}
