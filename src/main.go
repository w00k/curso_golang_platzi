package main

import "fmt"

// interfaz que ejecuta a la función area
type figura2D interface {
	area() float64
}

// cuadrado struct que contiene el lado del cuadrado
type cuadrado struct {
	base float64
}

// area implementación de la función de la interface
func (c cuadrado) area() float64 {
	return c.base * c.base
}

// rectangulo struct que contiene los lados del rectangulo
type rectangulo struct {
	base   float64
	altura float64
}

// area implementación de la función de la interface
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figura2D) {
	fmt.Println("area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	// con interfaz
	calcular(myCuadrado)
	calcular(myRectangulo)

	// lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
