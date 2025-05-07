package main

import (
	"fmt"
	"time"
)

func main() {
	go printHi()
	go fmt.Println("Hello from Main 2!")
	go fmt.Println("Hello from Main!")
	time.Sleep(time.Second)
}

func printHi() {
	fmt.Println("Hello from Goruntine!")
}
