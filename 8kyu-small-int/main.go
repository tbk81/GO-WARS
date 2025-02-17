package main

import "fmt"

func smallestInt(xi []int) (small int) {
	// small = 0
	for _, num := range xi {
		if num < small {
			small = num
		}
	}
	return
}

func main() {
	fmt.Println(smallestInt([]int{34, 15, 88, 2}))
}

/*
Given an array of integers your solution should find the smallest integer.
For example:

    Given [34, 15, 88, 2] your solution will return 2
    Given [34, -345, -1, 100] your solution will return -345
You can assume that the supplied array will not be empty.
*/
