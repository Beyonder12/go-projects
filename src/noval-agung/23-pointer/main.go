package main

import "fmt"

func main() {
	var a = 4
	var b *int = &a

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(*b)

	fmt.Println(&a)
	fmt.Println(&b)
}
