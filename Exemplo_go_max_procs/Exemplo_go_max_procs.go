package main

import (
	"fmt"
	"runtime"
	"time"
)

func mult_loop(a, b int) {
	c := 0
	for true {
		c = a * b
	}
	fmt.Println(c)
}

func main() {
	runtime.GOMAXPROCS(1)
	go mult_loop(10, 50)
	go mult_loop(15, 25)
	time.Sleep(time.Second * 100)
}
