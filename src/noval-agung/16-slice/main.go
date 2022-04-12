package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	bFruits[1] = "pinnaple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(string('c'), cap(baFruits))

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dst1 := []string{"potato", "potato", "potato"}
	src1 := []string{"watermelon", "pinnaple", "rere"}
	n1 := copy(dst1, src1)

	fmt.Println(dst1) // watermelon pinnaple potato
	fmt.Println(src1) // watermelon pinnaple
	fmt.Println(n1)   // 2

	var fruitss = make([]int, 5)
	fmt.Println(fruitss)
}

// var fruits = []string{"apple", "grape", "banana", "melon"}

// var aFruits = fruits[0:3]
// var bFruits = fruits[1:4]

// var aaFruits = aFruits[1:2]
// var baFruits = bFruits[0:1]

// fmt.Println(fruits)   // [apple grape banana melon]
// fmt.Println(aFruits)  // [apple grape banana]
// fmt.Println(bFruits)  // [grape banana melon]
// fmt.Println(aaFruits) // [grape]
// fmt.Println(baFruits) // [grape]

// // Buah "grape" diubah menjadi "pinnaple"
// baFruits[0] = "pinnaple"

// fmt.Println(fruits)   // [apple pinnaple banana melon]
// fmt.Println(aFruits)  // [apple pinnaple banana]
// fmt.Println(bFruits)  // [pinnaple banana melon]
// fmt.Println(aaFruits) // [pinnaple]
// fmt.Println(baFruits) // [pinnaple]
