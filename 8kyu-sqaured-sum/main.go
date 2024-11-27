package main

import "fmt"

func squareSum(xi []int) (sum int) {
	for _, num := range xi {
		sum = (num * num) + sum
	}
	return
}

func main() {
	fmt.Println(squareSum([]int{4, 10, -5, 100, 50, -10}))
}

/*
Complete the square sum function so that it squares each number passed into it and then sums the results together.
For example, for [1, 2, 2] it should return 9 because 12+22+22=91^2 + 2^2 + 2^2 = 912+22+22=9.
*/
