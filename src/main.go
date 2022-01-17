package main

import "fmt"

func main() {
	// llamamos una funcion
	normalFunction("Hola mundo")

	// llamamos una funcion con 3 parametos o argumentos
	tripleArgumentos(1, 2, "hola")

	// llamamos una funcion que retorna un numero
	fmt.Println(returnValue(2))

	// invocamos una funcion que retorna 2 valores
	value1, value2 := dolubleReturn(2)
	fmt.Println("value1:", value1, "\nvalue2:", value2)

	// invocamos la misma funcion que retorna 2 valores
	// y solo rescatamos uno de los valores (por orden)
	_, value := dolubleReturn(3)
	fmt.Println(value)
}

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgumentos(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dolubleReturn(a int) (c, d int) {
	return a, a * 2
}
