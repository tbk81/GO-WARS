package main

import (
	"fmt"
)

// func removeChar(word string) string {
// 	s := word[1 : len(word)-1]
// 	return s

func removeChar(word string) string {
	return word[1 : len(word)-1]
}

func main() {
	fmt.Println(removeChar("char"))
}

/*
create a function that removes the first and last characters of a string.
You're given one parameter, the original string. You don't have to worry about
strings with less than two characters.
*/
