package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Escrevendo 1...") // goroutine
	escrever("Escrevendo 2...")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
