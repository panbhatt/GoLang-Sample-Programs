package main

import (
	"fmt"
	"os"
)

func main(){
	var sum float64
	var n int
	var val float64
	for {
		if _,err :=fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
			fmt.Println("Exiting OS")
			os.Exit(1)
		}

	fmt.Println("FInal sum = " , sum)
}