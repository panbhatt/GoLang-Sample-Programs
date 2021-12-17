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

	a, b := int(10), uint32(64000)
	fmt.Printf("\n%10x %16b %32d", a,a,a)
	fmt.Printf("\n%10x %16b %32d", b,b,b)
}