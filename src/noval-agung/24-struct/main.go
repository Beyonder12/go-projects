package main

import "fmt"

func main() {
	// type student struct {
	// 	name  string
	// 	grade int
	// }

	// var s1 student
	// s1.name = "Fajri"
	// s1.grade = 100

	// fmt.Println(s1)

	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name  :", s1)
	fmt.Println("age   :", s1.age)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("grade :", s1.grade)

}

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}
