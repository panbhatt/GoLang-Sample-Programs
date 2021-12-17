package main

import (
	"fmt"
)

func sum(x int, nums ...int) int {
	sumf := x

	for _, num := range nums {
		sumf +=num
	}

	return sumf
}

func main() {

	fmt.Println(sum(1))
	fmt.Println(sum(1,2,3,4))

	sl := []int{1,2,3,4,5}
	fmt.Println(sum(5, sl...))


	fmt.Printf("%8b \n", 10)
	fmt.Printf("%4x \n", 10)
}