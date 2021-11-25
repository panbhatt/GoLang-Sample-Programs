package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"io"
	"bytes"
)

func main(){

	rdr := strings.NewReader("Pankaj Bhatt")

	/*
	if _, err := io.Copy(os.Stdout, rdr); err != nil {
		log.Fatal("AN Error occured, while copying")
		os.Exit(-1)
	} */


	var bf bytes.Buffer
	if _, err := io.Copy(&bf, rdr); err != nil {
		log.Fatal("AN Error occured, while copying")
		os.Exit(-1)
	}
	fmt.Println(bf.String())

	fmt.Println("DONE")



}