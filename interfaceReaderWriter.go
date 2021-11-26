// Program to show we can use Reader/Writer tCustom Implementation to 
// pass to the internal functions. 

package main

import (
	"fmt"
	"os"
	"io"
)

func main(){

	// OLD FUCNTION
	//func1()

	// New function
	func2()
}

func func1() {
	f1, _ := os.Open("./sample.txt")
	f2, _ := os.Create("./samplecopy.txt")

	defer f1.Close()
	defer f2.Close()

	n , _:=io.Copy(f2, f1)

	fmt.Printf("Func1() %d bytes copied \n", n)


}


type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return int(*b), nil
}

// Same as that of func1 but with own implementation.
func func2(){
	
	f1, _ := os.Open("./sample.txt")
	var f2 ByteCounter 

	defer f1.Close()

	n , _:=io.Copy(&f2, f1)

	fmt.Printf("Func2() %d bytes copied \n", n)
	


}