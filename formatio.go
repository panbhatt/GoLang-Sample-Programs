package main

import (
	"fmt"
)


func main(){


	a, b := 10,20
	c,d := 1.5, 3.23

	fmt.Printf("%d %d\n", a,b)
    fmt.Printf("%x %x\n", a,b)
	fmt.Printf("%#X %#X\n", a,b)

	fmt.Printf("%f %f\n", c,d)
	fmt.Printf("%6.2f %6.2f\n", c,d)

	// for slices. 
	fmt.Println("************** SLICES ************")
	s:= []int{1,2,3}


	fmt.Printf("%T \n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

	fmt.Println("************** ARRAY ************")
	ar:= [...]rune{'a','b','c'}


	fmt.Printf("%T \n", ar)
	fmt.Printf("%q \n", ar)
	fmt.Printf("%v\n", ar)
	fmt.Printf("%#v\n", ar)

	fmt.Println("************** Map ************")
	mp:= map[string]int{ "a" : 1, "b" : 2}


	fmt.Printf("%T \n", mp)
	fmt.Printf("%v\n", mp)
	fmt.Printf("%#v\n", mp)

	fmt.Println("************** STRING ************")
	str :="Pankaj Bhatt"


	fmt.Printf("%T \n", str)
	fmt.Printf("%q \n", str)
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)

	fmt.Println("************** BYTES ************")
	bs :=[]byte(str)


	fmt.Printf("%T \n", bs)
	fmt.Printf("%q \n", bs)
	fmt.Printf("%v\n", bs)
	fmt.Printf("%#v\n", bs)

}