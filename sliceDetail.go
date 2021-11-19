package main

// Various wake of making slices . 

import (
	"fmt"
	"encoding/json"
)

func main() {


	var s []int

	fmt.Printf("%d %d %T %5t, %#[3]v \n", len(s), cap(s), s, s==nil )
	sj, _ := json.Marshal(s)
	fmt.Println("JSON => ", string(sj))

	s1 := []int{}
	fmt.Printf("%d %d %T %5t, %#[3]v \n", len(s1), cap(s1), s1, s1==nil )
	sj1, _ := json.Marshal(s1)
	fmt.Println("JSON => ", string(sj1))

	// Append to this, will add the element at 6 position. 
	s2 := make([]int, 5) // len
	fmt.Printf("%d %d %T %5t, %#[3]v \n", len(s2), cap(s2), s2, s2==nil )
	sj2, _ := json.Marshal(s2)
	fmt.Println("JSON => ", string(sj2))

	s3 := make([]int, 0, 5) // len
	fmt.Printf("%d %d %T %5t, %#[3]v \n", len(s3), cap(s3), s3, s3==nil )
	sj3, _ := json.Marshal(s3)
	fmt.Println("JSON => ", string(sj3))

	// Check for SLICE Emptiness
	if len(s) == 0 {
		fmt.Println("Length of S nil slize is zero")
	}

	s = append(s,10)
	fmt.Println(s)

	newSl := s[0:1:1] // means len and cap would be 1
	fmt.Printf("LEN = %d CAP = %d " , len(newSl), cap(newSl))

}