package main

import "fmt"

func main() {
	//for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(" ")

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println(" ")

	// for forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
