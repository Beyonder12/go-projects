package main

import "fmt"

func sendValue(c chan int) {
	c <- 8
}

func main() {
	fmt.Println("Go Channel Tutorial")

	values := make(chan int)
	defer close(values)

	go sendValue(values)

	value := <-values
	fmt.Println(value)
}
