package main

import (
	"fmt"
)

func vowelCount(str string) int {
	vowel := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	total := 0
	for _, char := range str {
		if vowel[char] {
			total++
		}
	}
	return total

	// A much better solution using case switching the other solution was from claude
	// for _, c := range str {
	// 	switch c {
	// 	case 'a', 'e', 'i', 'o', 'u':
	// 	  count++
	// 	}
	//   }
	//   return count
	// }
}

func main() {
	fmt.Println(vowelCount("Hello"))
}

/*
Return the number (count) of vowels in the given string.
We will consider a, e, i, o, u as vowels for this Kata (but not y).
The input string will only consist of lower case letters and/or spaces.

*/
