package main

import "fmt"

func main() {
	// constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi = ", pi)
	fmt.Println("pi2 = ", pi2)

	// variables
	base := 12          //no le decimos el tipo de dato, pero si el valor
	var altura int = 14 //definimos el tipo de dato y valor
	var area int        //definimos el tipo de dato

	fmt.Println(base, altura, area)

	// zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	// area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("area cuadrado:", areaCuadrado)
}
