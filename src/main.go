package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"hola", "que", "tal", "123"}

	// con indice
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	fmt.Println("")

	// sin indice
	for _, valor := range slice {
		fmt.Println(valor)
	}

	fmt.Println("")

	// solo los indices
	for i := range slice {
		fmt.Println(i)
	}

	isPalindromom("Ama")
	isPalindromom("esto no es palindromo")
	isPalindromom("amor a Roma")
}

func isPalindromom(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	fmt.Println("")
	fmt.Println("el texto inrgesado es '", text, "'")
	fmt.Println("el texto al revez es '", textReverse, "'")

	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("es un palindromo")
	} else {
		fmt.Println("no es un palindromo")
	}
}
