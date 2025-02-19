package main

import (
	"fmt"
	"strings"
)

func accum(s string) string {
	var word string
	xs := strings.Split(s, "")
	for i := 0; i < len(xs); i++ {
		word += xs[0]
	}
	return word
}

func main() {
	ans := accum("hello")
	fmt.Println(ans)
}

/*
Examples:

accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
*/
