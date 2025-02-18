package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Escrevendo 1..."
	canal <- "Escrevendo 2..."

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
