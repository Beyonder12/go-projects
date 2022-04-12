package main

import "fmt"

func main() {
	// str := "abc"
	// chars := []rune(str)
	// for i := 0; i < len(chars); i++ {
	// 	char := string(chars[i])
	// 	println(char)
	// }
	main1()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	// add(20, 30)

}

// func getLucky(s string, k int) int {
// 	var sTr := convert(s);
// 	var temp := sTr;
// 	var result = "";
// }

// func convert(s string) string {

// 	result
// }

// Function accepting arguments
// func add(x int, y int) {
// 	total := 0
// 	total = x + y
// 	fmt.Println(total)
// }

// class Solution {
//     public int getLucky(String s, int k) {
//         String sTr = convert(s), temp = sTr, result = "";

//         for(int i = 0 ; i < k; i++) {
//             result = transform(temp);
//             temp = result;
//         }

//         return Integer.parseInt(result);

//     }

//     private String transform(String s) {
//         int result = 0;
//         for(char ch : s.toCharArray()) {
//             result += (ch % 48);
//         }

//         return result + "";
//     }

//     private String convert(String s) {
//         String result = "";
//         for(char ch: s.toCharArray()) {
//             result += (ch % 96);
//         }
//         return result;
//     }
// }
