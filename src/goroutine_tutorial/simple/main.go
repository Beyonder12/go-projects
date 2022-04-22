package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concurrency With Goroutine")
	go calculate(4)
	calculate(5)
	// go calculate(6)

	fmt.Scanln()

}

func calculate(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
