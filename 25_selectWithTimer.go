/**
This program demonstrate how we can apply timeouts via using
timer
*/

package main

import (
	"log"
	"time"
)

func main(){

	const tickRate = 2 * time.Second

	stopper := time.After(5 *tickRate)
	ticker := time.NewTicker(tickRate).C 
	
	loop:
		for {
			select {
			case val:= <-ticker:
				log.Println("Ticker -> " , val)
			
			case t := <- stopper:
				log.Println(" Stopped -> ", t)
				break loop
			
			// You can write a default case, that will get executed
			// if none of the other cases gest executed, but of no
			// use as of now
			/*
			default :
				log.Println("Default")
			*/

			}
		}
	

}