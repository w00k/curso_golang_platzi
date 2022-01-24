package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")

	wg.Add(1)
	// concurrencia con goroutines
	go say("world", &wg)
	wg.Wait()

	// funciones anónimas
	go func(text string) {
		fmt.Println(text)
	}("adios")
	// agrego tiempo de espera para que se imprime la goroutines
	time.Sleep(time.Second * 1)
}
