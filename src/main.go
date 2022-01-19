package main

import "fmt"

func main() {
	// array, el largo es inmutable
	var array [4]int
	array[0] = 1
	array[1] = 2

	fmt.Println(array)
	fmt.Println("largo del array ", len(array))
	fmt.Println("capacidad de la array ", cap(array))

	// slice, son mutables
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(slice)
	fmt.Println("largo del slice ", len(slice))
	fmt.Println("capacidad del slice ", cap(slice))

	// metodos en el slice
	fmt.Println(slice[0])   //imprime el elemento de la posicion 0
	fmt.Println(slice[:3])  //imprime hasta la posicion 3, posicion 3 es exclusivo
	fmt.Println(slice[5:8]) //imprime entre la posicion 5 y 8, aca las posiciones son inclusivas
	fmt.Println(slice[4:])  //imprime desde el elemento 4 en adelante, posicion exclusivo

	// agregar elementos al slice
	slice = append(slice, 10)
	fmt.Println(slice)

	newSlice := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
