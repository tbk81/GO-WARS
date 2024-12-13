package main

import "fmt"

func ghSum(num int) int {
	var total int
	for i := 1; i < num+1; i++ {
		total += i
		// fmt.Println(total)
	}
	return total
	// Do this in one line!
	// return n * (n + 1) / 2
}

func main() {
	fmt.Println(ghSum(5))
}

/*
Write a program that finds the summation of every number from 1 to num.
The number will always be a positive integer greater than 0.
Your function only needs to return the result, what is shown between
parentheses in the example below is how you reach that result and it's
not part of it, see the sample tests.
*/
