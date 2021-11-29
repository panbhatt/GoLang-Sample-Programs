// Class 19

package main


import (
	"fmt"
)

type Pair struct {
	Path string 
	Hash string 
}

func (p Pair) String() string {
	return fmt.Sprintf(" Hash of %s is %s ", p.Path , p.Hash)
}


type PairWithLength struct {
	Pair // This is called promotion. even if it is an Pointer it will work. 
	Length int
}

func (pl PairWithLength) String() string {
	return fmt.Sprintf(" Hash of %s is %s and Length = %d ", pl.Path , pl.Hash, pl.Length)
}

func main() {

	p := Pair{"Pankaj Bhatt", "0xfdfe"}

	fmt.Println(p)

	pl := PairWithLength{ Pair{"ram", "0xdead"}, 5}

	fmt.Println(pl)  // will call Pair string method unless you have
	// override the function in the derived.


	// Pair and PairWithLength are not same. 
	printPath(p) // will work
	// printPath(pl) // will nto work
	printPath(pl.Pair) // will work.

}

func printPath(p Pair)  {
	 fmt.Println(p.Path)
}
