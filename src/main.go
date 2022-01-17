package main

import "fmt"

func main() {

	// area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// suma
	result := x + y
	fmt.Println("suma ", result)

	// resta
	result = y - x
	fmt.Println("resta ", result)

	// multiplicacion
	result = x * y
	fmt.Println("multiplicacion ", result)

	// division
	result = y / x
	fmt.Println("division ", result)

	// modulo
	result = y % x
	fmt.Println("modulo ", result)

	// incremental
	// x + 1
	x++
	fmt.Println("x incremental ", x)

	// decremental
	// x - 1
	x--
	fmt.Println("x decremental ", x)

	// reto
	// calcular area de un trapecio
	var base1 = 4
	var base2 = 10
	var altura int = 4
	fmt.Println("area del trapecio es ", areaTrtapecio(base1, base2, altura))

	// calcular area de una circunferencia
	var radio int = 5
	fmt.Println("area de la circunferencia ", 5, " es ", areaCirculo(radio))

}

func areaTrtapecio(base1 int, base2 int, altura int) int {
	return ((base1 + base2) * altura) / 2
}

func areaCirculo(radio int) float64 {
	const pi float64 = 3.14
	return pi * float64(radio) * float64(radio)
}
