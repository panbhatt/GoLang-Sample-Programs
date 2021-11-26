package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

// You can put method on any USER DECLARED TYPE... not on defined types already. 

func (is IntSlice) String() string{

var str []string

for index,i := range is {
	str = append(str, strconv.Itoa(index) + ": " +strconv.Itoa(i) + " ")
}

return "[" + strings.Join(str,",") + "]"

}

func main() {
	var is IntSlice = []int{100,200,3,4,5}
	var s fmt.Stringer // its an interface variable.

	fmt.Println(is)
	fmt.Printf("%T  %[1]v \n", is)

	s = is

	fmt.Printf("%T %[1]v\n", s)
}