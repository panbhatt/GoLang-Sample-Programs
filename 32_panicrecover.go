package main

import (
	"fmt"
)

func main() {
	fmt.Println("Before Gen Panic")
	defer func () {

		if p:= recover(); p!= nil {
			fmt.Println(" I m trying to recover.")
			fmt.Println("Recover : " , p)
		}
	}()
	genPanic()
}

func genPanic(){
	panic("OMG something hapeend. ")
}