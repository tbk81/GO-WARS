package main

import (
	"fmt"
	"strconv"
)

func IntToStr(x int) string {
	return strconv.Itoa(x)
}

func main() {
	x := 42
	fmt.Printf("Before func: %v is of type %T\nAfter func: %v is of type %T\n", x, x, IntToStr(x), IntToStr(x))
}

/*
We need a function that can transform a number (integer) into a string.
*/
