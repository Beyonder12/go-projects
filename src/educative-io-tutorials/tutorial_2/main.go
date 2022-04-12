package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable1 string
	var variable2 int
	var variable3 float64
	var variable4 bool
	var variable5 = []string{}

	fmt.Println("Type of variable1: ", reflect.ValueOf((variable1)).Kind())
	fmt.Println("Type of variable2: ", reflect.ValueOf((variable2)).Kind())
	fmt.Println("Type of variable3: ", reflect.ValueOf((variable3)).Kind())
	fmt.Println("Type of variable4: ", reflect.ValueOf((variable4)).Kind())
	fmt.Println("Type of variable5: ", reflect.ValueOf((variable5)).Kind())

	fmt.Printf("Type of variable1: %T\n", variable1)
	fmt.Printf("Type of variable2: %T\n", variable2)
	fmt.Printf("Type of variable3: %T\n", variable3)
	fmt.Printf("Type of variable4: %T\n", variable4)
	fmt.Printf("Type of variable5: %T\n", variable5)

}
