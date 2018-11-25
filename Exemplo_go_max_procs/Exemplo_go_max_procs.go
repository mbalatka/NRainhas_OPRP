package main

import (
	"fmt"
	"runtime"
	"time"
)

func mult_loop(a, b int) {
	c := 0
	for true {
		c = a * b //simulacao de processamento
	}
	fmt.Println(c)
}

func main() {
	runtime.GOMAXPROCS(1) //define o numero maximo de processadores utilizados

	go mult_loop(10, 50) //chamas as goroutines
	go mult_loop(15, 25)

	time.Sleep(time.Second * 100)
}
