package main

import "fmt"

func strRepeat(reps int, str string) string {
	for n := 1; n < reps; n++ {
		str += str
	}
	return str
}

func main() {
	fmt.Println(strRepeat(4, "Gopher"))
	// str := "hello"
	// count := 6
	// for n := 0; n < count; n++ {
	// 	fmt.Print(str)
	// }
	// fmt.Println()
}

/*
Write a function that accepts an integer n and a string s as parameters,
and returns a string of s repeated exactly n times.
*/
