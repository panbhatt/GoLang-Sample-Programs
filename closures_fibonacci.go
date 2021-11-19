package main

import (
	"fmt"
)

func main(){

	fmt.Printf("**** FIBONACCI SERIES ********\n")

	fmt.Printf("%d %d", 0,1)
	fibClosure := getFibClosure()

	for i:=0;i<20;i++ {
		fmt.Printf(" %d ", fibClosure())  // calling a closure. 
	}


	fmt.Println("\n SEcond Series ");
	fib2ndClosure := getFibClosure()

	for i:=0;i<5;i++ {
		fmt.Printf(" %d", fib2ndClosure())
	}


}

func getFibClosure() func() int{
	a, b := 0,1

	// This is the closure. 
	fibClosure := func() int {
		a, b = b, a+b 
		return b
	}

	return fibClosure
}