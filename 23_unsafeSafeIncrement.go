package main

import ( "fmt"
	"net/http"
	"log"
)

var nextInt = 0; 


// This is an Unsafe Handler, means, multiple hthreads can modify the same value
// and return it. 
func UnsafeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Next Id is %d", nextInt )
	nextInt++
}

// This is a safe way to do the increments. 

type NextIntChan chan int  // doing OOP of INT CHANNEL

func (n NextIntChan) increment(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Safe Next Int is %d ", <-n) // Reading from Int Chan

}

func increment(niCh chan<- int) {
	for i:=1 ; ; i++ {
		niCh<- i
	}
}

func main() {


	

	// Below two lines is for safe execution. THIs is extra imporatnt to use the interfaces
	var niCh NextIntChan = make(chan int)

	go increment(niCh)

	http.HandleFunc("/inc/unsafe", UnsafeHandler)
	http.HandleFunc("/inc/safe", niCh.increment)  // use watch -n 0 to repeat request every second. 

	log.Fatal(http.ListenAndServe(":8000", nil))
	log.Printf("Server Started at 8000. \n")

}



