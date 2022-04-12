package main

import "fmt"

func main() {
	var avg = calculate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Printf("Average : %.2f\n", avg)

}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return float64(total) / float64(len(numbers))
}
