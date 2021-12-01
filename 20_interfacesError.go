package main

import (
	"fmt"
	
)

type CustomError struct {
		err error
		path string
}

func (er CustomError) Error() string {
	return fmt.Sprintf("Err -> %s , path -> %s", er.err, er.path)
}

func XYZ(a int) *CustomError {
	return nil
}


func XYZ2(a int) error { // THis is the corrrect WAY.
	return nil
}

// Main Fucntion

func main(){

	var er error = XYZ(1)


	// this is amazing, it has to be NIL, but remember class 20, every interface
	// has two parts, TYPE it points to (which is not nil) and the pointer to the acutal error
	if er != nil {
		fmt.Println("ERR is not null, did something happened")
	} else {
		fmt.Println("Err is NILL")
	}

	fmt.Println("******************")
	var er1 error = XYZ2(1)


	// this is amazing, it has to be NIL, but remember class 20, every interface
	// has two parts, TYPE it points to (which is not nil) and the pointer to the acutal error
	if er1 != nil {
		fmt.Println("ERR is not null, did something happened")
	} else {
		fmt.Println("Err is NILL")
	}


}