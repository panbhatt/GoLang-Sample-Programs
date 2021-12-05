// This function solve the probelm of genreating the PRime no but with the help fo channels. 
// Gnerator will take a limit and applky a filter and return a new channel if it passes again 
// then it will continue till we reaches n/2. Please see this https://www.youtube.com/watch?v=zJd7Dvg3XCk&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=24


package main

import (
	"fmt"
)

/**
	limit : it refers to the 
*/
func generator(limit int, ch chan<- int) {

	for i:=2  ;i < limit ; i++{

		ch <- i 
	}

	close(ch)

}

func filter(src <-chan int, dst chan<- int, prime int) {

	for i:= range src{

			if i%prime != 0 {
				dst <- i
			}

	}

	close(dst)
}

func main() {
	sieve(100)
}

func sieve(limit int){

	ch := make(chan int)
	go generator(limit, ch)

	for {

		prime, ok := <-ch
		if !ok {
			break	
		}

		ch1 := make(chan int)
		go filter(ch,ch1,prime)

		ch = ch1
		fmt.Print(prime, " " )
	}

	fmt.Println()

}

