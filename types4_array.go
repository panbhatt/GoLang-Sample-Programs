package main

import (
	"fmt"
)

func main() {

	// Create Arrayls. 
	
	var t = []byte("string")

	fmt.Println(len(t), t)
	fmt.Println(t[1:])
	fmt.Println(t[:5])

	var s = [3]int{1,2,3} 
	fmt.Println(s)
	s = [...]int{10,20,30}
	fmt.Println(s)
	s1 := make([]int, 5)

	 s2 :=[]int{100,200,2000}
	copy(s1,s2)  // both these needs to be slice
	fmt.Println(s1)
	fmt.Println(">> Logging the fucntion call paramters")
	justPrint(s,s2)
}

func justPrint(sar [3]int, ssl []int ) {
	fmt.Println(sar)
	fmt.Println(ssl)
}
