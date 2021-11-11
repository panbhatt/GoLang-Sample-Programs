package main

import (
	"fmt"
)

func main(){
	var a int = 10 
	var (
		i = 10
		j = 10.23
	)

	
	fmt.Printf("%T %v\n", j, j)
	fmt.Printf("%T %[1]v\n", a)
	fmt.Println(i,j,a)

	a=int(i)
	j = float64(10.23)

	fmt.Printf("%T   %[1]v\n", j )

}