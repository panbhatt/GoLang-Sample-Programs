/**
Print list of wrds characters and lines in each and every file that is passed as an argument. 

*/

package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	"bufio"
	"strings"
)

func main(){
	 
	listOfFiles := os.Args[1:]

	if len(listOfFiles) > 0 {
		for _,fName := range(listOfFiles) {
		fl, err := os.Open(fName)
	

		if err == nil {
			fmt.Printf("\n**%s", fName)
			/*if data,er:= ioutil.ReadAll(fl); er == nil {
				
				fmt.Printf("%s",data);
				fmt.Println("*** Size = "+ len(data))
				
			} */

			countDetails(fl)
			
		} else {
			fmt.Println("An Error occured, while reading file ", fName)
		}

		defer fl.Close()
		}

	} else {
		fmt.Fprintln(os.Stderr, " Error - No input files passed. ");
	}



}

func countDetails(fl *os.File){

	scan :=  bufio.NewScanner(fl);
	var lc, wc, cc int

	for scan.Scan(){

		s := scan.Text();
		lc++
		wc += len(strings.Fields(s)		)
		cc += len(s)
	}

	fmt.Printf("%5d %5d %5d\n", lc, wc,cc)

}
