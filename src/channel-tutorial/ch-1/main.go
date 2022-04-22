package main

import "fmt"

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	v1 := <-fooVal
	v2 := <-fooVal

	fmt.Println(v1, v2)

	// arr := []int{5, 4, 3, 4, 5}
	str := "abcde"
	fmt.Println(isPalindrome(str))
	fmt.Println(reverseString(str))

}

func reverseString(str string) string {
	byte_str := []rune("asdfg")
	fmt.Println(byte_str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	fmt.Println(string(byte_str))
	return string(byte_str)
}

func isPalindrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
