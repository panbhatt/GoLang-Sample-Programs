package main

import (
	"fmt"
)

func main(){

	
	fmt.Println("************ ARRAY *********")
	ar := [3]int{1,2,3}

	doArray(ar)
	fmt.Printf("%#v\n", ar)
	fmt.Println("************ SLICES *********")

	sl := []int{1,2,3}
	doSlice(sl)
	fmt.Printf("%#v \n", sl)

	fmt.Println("************ MAPS *********")
	mp := map[string]int{
		"ram" : 12,
		"shyam" : 14, 
	}
	doMap(mp)
	fmt.Printf(" %#v\n ", mp)

	fmt.Println("************ MAPS BY ADDRESS *********")

	doMapByAddress(&mp)
	fmt.Printf(" IN main function map -> %#v \n", mp)
	

}

// Passed by Value
func doArray(b [3]int) {
	b[0]=100
	fmt.Printf("In function doArray -> %#v\n", b)
}

// Passed by reference
func doSlice(b []int) {
	b[0]=200
	fmt.Printf("In Function doSlice -> %#v %p   \n" , b, b)
}

// Passed by refrence
func doMap(mp map[string]int) {
		mp["ram"] = 13344
		fmt.Printf(" IN function doMap -> %#v %p  \n", mp, mp)
}

func doMapByAddress(mp *map[string]int){
	
	fmt.Printf(" In doMayByAddress -> %#v \n ", mp)
	fmt.Println("Ram value is " , (*mp)["ram"])

	

	// or the above statement can be
	*mp = map[string]int{
		"canada" : 200, 
		"usa" : 234,
	}

	fmt.Printf("Changed value -> %#v \n ", mp)

}