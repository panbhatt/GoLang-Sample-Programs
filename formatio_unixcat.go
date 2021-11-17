/**
Implementation CAT

*/

package main

import (
	"fmt"
	"os"
	"io"
)

func main(){
	 
	listOfFiles := os.Args[1:]

	if len(listOfFiles) > 0 {
		for _,fName := range(listOfFiles) {
		fl, err := os.Open(fName)

		if err == nil {
			fmt.Printf("\n**%s\n", fName)
			if _,er:= io.Copy(os.Stdout, fl); er != nil {
				fmt.Println("Error occured, while rading file ", fName)

				
			}
			
		} else {
			fmt.Println("An Error occured, while reading file ", fName)
		}

		defer fl.Close()
		}

	} else {
		fmt.Fprintln(os.Stderr, " Error - No input files passed. ");
	}



}
