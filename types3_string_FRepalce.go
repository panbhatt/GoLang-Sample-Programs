package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Wrong number of arguments passed.")
		os.Exit(1)
	}

	oldWord , newWord := os.Args[1], os.Args[2]
	
	// now start taking input the data. 
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(),oldWord)
		t := strings.Join(s,newWord)
		fmt.Println(t)
	}

}