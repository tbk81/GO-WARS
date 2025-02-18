package main

import "fmt"

func smallestInt(xi []int) (small int) {
	small = xi[0]
	for _, num := range xi {
		if num < small {
			small = num
		}
	}
	return
}

func largesttInt(xi []int) (large int) {
	for _, num := range xi {
		if num > large {
			large = num
		}
	}
	return
}

func main() {
	fmt.Println(smallestInt([]int{34, 15, 88, 2}))
	fmt.Println(largesttInt([]int{34, 15, 88, -2, 100, -90}))
}

/*
In this little assignment you are given a string of space separated numbers,
and have to return the highest and lowest number.
Examples

HighAndLow("1 2 3 4 5") // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"

*/
