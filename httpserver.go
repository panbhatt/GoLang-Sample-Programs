package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintf(rw,"hello world from Http Server %s \n", r.URL.Path[:])
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))


}



