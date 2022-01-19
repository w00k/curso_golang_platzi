package main

import "fmt"

func main() {
	defer fmt.Println("hola") // defer se ejecuta al final, por eso impime al final "hola"
	fmt.Println("mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		if i == 8 {
			fmt.Println("break")
			break
		}
	}
}
