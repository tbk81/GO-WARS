package main

import "fmt"

func SumPos(xs []int) (sum int) {
	for _, num := range xs {
		if num > 0 {
			sum = sum + num
		}
	}
	return
}

func main() {
	fmt.Println(SumPos([]int{2, 6, -5, -10, 11}))
}

/*
In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?
The number can be negative already, in which case no change is required.
Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.
*/
