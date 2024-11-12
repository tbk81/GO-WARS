package main

import "fmt"

func ReverseString(str string) string {
	var result = ""
	for _, char := range str {
		result = string(char) // + result
		return result
	}
	// return result
}

func main() {
	fmt.Println(ReverseString("Hello"))
	fmt.Println(ReverseString("World!"))
}

/*
Complete the solution so that it reverses the string passed into it.
*/
