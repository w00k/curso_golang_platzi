package main

import "fmt"

func main() {

	isOdd(2)
	isOdd(5)

	fmt.Println()

	// sin condicion, para multiples variables
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}

func isOdd(number int) {
	//asi o del otro modo, anidado ene l switch
	//modulo := number % 2
	switch modulo := number % 2; modulo {
	case 0:
		fmt.Println("false, numero ", number, " es par")
	default:
		fmt.Println("true, numero ", number, " es impar")
	}
}
