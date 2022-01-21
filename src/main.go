package main

import (
	"fmt"

	pk "github.com/w00k/curso_golang_platzi/src/mypackage"
)

func main() {
	myCar := pk.CarPublic{
		Brand: "Ford",
		Year:  2022,
	}

	fmt.Println(myCar)

	var otherCar pk.CarPublic
	otherCar.Brand = "Ferrari"

	fmt.Println(otherCar)
}
