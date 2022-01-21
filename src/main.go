package main

import (
	"fmt"
)

// pc es un struct para el ejemplo
type pc struct {
	ram   int
	disk  int
	brand string
}

// ping método de struc 'pc', que imprime la marca
func (myPc pc) ping() {
	fmt.Println(myPc.brand, " Pong")
}

// duplicateRAM es una función que permite acceder al atributo 'ram' y multiplicarlo por 2
// usando punteros, mejor practica
func (myPc *pc) duclicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(a)  // valor
	fmt.Println(b)  // dirección de memoria
	fmt.Println(*b) // valor de la dirección de memoria

	*b = 100 // nuevo valor a la dirección de memoria

	fmt.Println(a) // valor, a y *b están apuntando a la misma dirección de memoria

	myPc := pc{
		ram:   16,
		disk:  200,
		brand: "msi",
	}

	fmt.Println(myPc)
	myPc.ping()

	//myPc.ram = 32 // puede ser rebundante

	myPc.duclicateRAM() // convención para modificar valores de atributos
	fmt.Println(myPc)
}
