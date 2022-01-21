package main

import (
	"fmt"
)

func main() {
	// es mas eficiente que usar slice, usa concurrencia por lo que
	// los valores puden no recorrerse en orden
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// un valor
	fmt.Println(m["Jose"])
	fmt.Println(m["Pepito"])
	fmt.Println(m["no existe"]) // es cero

	//verificar si excte en el diccionario
	value, ok := m["Pepito"]
	fmt.Println(value, ok) // OK = true

	valueNotFount, nok := m["not found"]
	fmt.Println(valueNotFount, nok) // NOK = false

	// recorrer el map
	for i, valor := range m {
		fmt.Println(i, valor)
	}
}
