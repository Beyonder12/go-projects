package main

import (
	"fmt"
	"time"
)

func main() {
	count("--------------sheep")
	go count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 2000)
	}
}
