package main

import (
	"fmt"
	"strings"
)

// var vowels []string
func disemvowel(str string) string {
	// vowels = "a", "e", "i", "o", "u", "A", "E", "I", "O", "U"
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, v := range vowels {
		str = strings.ReplaceAll(str, v, "")
	}
	return str
}

func main() {
	fmt.Println(disemvowel("This website is for losers LOL!"))
}

/*
Your task is to write a function that takes a string and return a new string with all vowels removed.
For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".
Note: for this kata y isn't considered a vowel.
*/
