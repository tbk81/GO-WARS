package main

import (
	"fmt"
	"strings"
)

func getMidChar(s string) string {
	xs := strings.Split(s, "")
	ld2 := len(xs) / 2
	var newString string
	if (len(xs))%2 == 0 {
		newString = xs[ld2-1] + xs[ld2]
	} else {
		newString = xs[ld2]
	}
	return newString
}

func main() {
	i := "A"
	x := getMidChar(i)
	fmt.Println(x)
	fmt.Printf("len: %v\nlen/2: %v\n", len(i), len(i)/2)
	fmt.Println(len(i) % 2)

}

/*
You are going to be given a non-empty string. Your job is to return the middle character(s) of the string.

    If the string's length is odd, return the middle character.
    If the string's length is even, return the middle 2 characters.

Examples:

"test" --> "es"
"testing" --> "t"
"middle" --> "dd"
"A" --> "A"


n := len(s)
    if n==0 {return s}
    return s[(n - 1) / 2 : n / 2 + 1]


OR

mid := len(s)/2
  if len(s) % 2 == 1 {
    return string(s[mid])
  }
  return string(s[mid-1]) + string(s[mid])

*/
