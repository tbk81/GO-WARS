package main

import "fmt"

// var x int

func NegNum(x int) int {
	if x >= 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(NegNum(-600))
}

/*
You get an array of numbers, return the sum of all of the positives ones.
Example [1,-4,7,12] => 1 + 7 + 12 = 20
Note: if there is nothing to sum, the sum is default to 0.

*/
