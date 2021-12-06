// Introduction of select, so that you can listen on multiple 
// channels at the same time. 


package main

import (
	"log"
	"time"
)

func main() {

	// Creation of two channels
	chans := []chan int {
		make(chan int),
		make(chan int),
	}

	// Invoke two goroutine that publish on those channels. 

	for i:=1;i<=2;i++ {

		go writeNo(i, chans[i-1])
	
	}

	// Read from both othe channle only 10 no of times. 
	for i:=0;i<10;i++ {
		select {
		case val := <- chans[0]:
			log.Println("Channel 1 -> " , val)
		case val := <- chans[1]:
			log.Println("Channel 2 -> " , val)

		}
	}


}


func writeNo(n int, ch chan<- int) {

	for {
		time.Sleep(time.Duration(n) *  time.Second)
		ch <- n
	}

}
