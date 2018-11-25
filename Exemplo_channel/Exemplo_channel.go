package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process(a, b int, ch chan int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(4))) //sleep aleatorio ate 4 segundos
	result := a * b                                       //processa
	ch <- result                                          //envia o resultado para o channel
}

func main() {
	ch := make(chan int, 1) //cria o canal

	go process(10, 5, ch) //cria as go routines
	go process(100, 7, ch)
	go process(1000, 8, ch)

	somatorio := 0

	for i := 0; i < 3; i++ { //recebe todas as respostas do canal
		result := <-ch
		fmt.Println("Recebeu:", result)
		somatorio += result
	}
	fmt.Println("Somatorio:", somatorio) //exibe a resposta
}
