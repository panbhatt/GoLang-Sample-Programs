package main

import (
	"fmt"
)

func main(){

	var album = struct {
		title string
		artist string
	}{
		"While Album",
		"The Beatles",
	}

	var pAlbum = &album

	fmt.Println(album, pAlbum)



	// Anonymous structs can be assinged to each other. 
	s1 := struct{
		name string
	} { "Pan"}

	s2 :=  struct {
		name string
	} { "Rahul"}

	s1 = s2	

	fmt.Println(s1, s2)

	// If the struct is named it can't be assigned (only Explicitly)
	type Employee struct { name string }
	type Boss struct { name string }

	e1 := Employee { "Pankaj"}
	e2 := Boss { "Steve"}

	e1 = Employee(e2)

	fmt.Println(e1, e2)
	

}