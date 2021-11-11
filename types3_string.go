package main

import (
	"fmt"
	"strings"
)

func main(){
	var s string

	// string is immutable, u can't change any char of string.
	s = "Pankaj Bhatt ₹"

	fmt.Printf("%10T %[1]v\n", s)
	fmt.Printf("%10T %[1]v\n", []rune(s))
	fmt.Printf("%10T %[1]v\n", []byte(s))
	fmt.Println("%10T% %[1]v\n", s, len(s)) // len of the bytes.

	a := s[3:]
	b := s[:3]
	fmt.Println(a, b)

	var st string = s + " ; "
	fmt.Println(st)

	// String functions. there are tons. 
	fmt.Println(strings.Contains(s,"Pan"))
	fmt.Println(strings.HasPrefix(s ,"bsd" ))
	fmt.Println(strings.ToUpper(s))

}