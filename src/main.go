package main

import "fmt"

// pc es un struct para el ejemplo
type pc struct {
	ram   int
	disk  int
	brand string
}

// String personaliza el output del struct retorno del objeto ps retornando una cadena como string
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, brand: "msi", disk: 200}

	fmt.Println(myPc)
}
