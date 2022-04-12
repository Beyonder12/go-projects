package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(reverse(arr[:]))

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for key, fruit := range fruits {
		fmt.Printf("index -%d nama buah : %s\n", key, fruit)
	}

}

func reverse(arr []int) []int {

	for i, j := 0, len(arr)-1; i < j; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		fmt.Printf("The index %d : %d\n", i, arr[i])
		j--
	}
	return arr
}
