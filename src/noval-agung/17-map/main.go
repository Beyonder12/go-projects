package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["January"] = 50
	chicken["February"] = 4

	fmt.Println(chicken)

	chicken = map[string]int{
		"January":  10,
		"February": 20,
	}

	var key, value = chicken["January"]

	fmt.Println(key, value)

	delete(chicken, "January")
	fmt.Println(chicken)
}
