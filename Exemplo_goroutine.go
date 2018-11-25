package main

import (
	"fmt"
	"time"
)

func print_5x(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Println(s)
	}
}

func main() {
	print_5x("String 1")
	print_5x("String 2")
	time.Sleep(time.Second * 5)
}
