package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(10000)
		fmt.Println("Inside go routine")
		ch <- "Hello World"
	}()

	x := <-ch
	fmt.Println(x)
}
