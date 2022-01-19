package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// con AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("AND Es verdad")
	}

	// con OR
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("OR Es verdad")
	}

	// funcion que determine si un numero es par o impar
	isPair(1)
	isPair(2)
	isPair(17)
	isPair(18)
}

func isPair(number int) {
	if number%2 == 0 {
		fmt.Println("numero ", number, " es par")
	} else {
		fmt.Println("numero ", number, " es impar")
	}
}
