

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	

)

// This program will never print out 1000.
func main() {
	// This function will print the last value of the integter
	fmt.Println("Race Condition Increment = " , do())

	// This function will print the last value of the integter
	fmt.Println("With COunting Semaphore " , doWithCountingSemaphore())

	fmt.Println("With COunting Semaphore via MUTEX " , doWithMutex())

	fmt.Println("With COunting Semaphore via ATOMIC " , doWithSyncAtomic())

	
}

func do() int {
	var n int64

	var wg sync.WaitGroup

	for i:=0;i<1000;i++{

		go func() {
			wg.Add(1)
			n++  // THis is the function that would generate the race condition. 
			wg.Done()
		}()
	}

	wg.Wait()
	return int(n)
}


func doWithCountingSemaphore() int {

	var n int64
	var wg sync.WaitGroup
	m:=make(chan bool, 1)

	for i:=0 ;i <1000;i++ {
		wg.Add(1)
		go func() {
			m <- true
			n++
			<-m
			wg.Done()
		}()
	}

	wg.Wait()
	return int(n)

}


func doWithMutex() int {

	var n int64
	var mx sync.Mutex
	var wg sync.WaitGroup


	for i:=0 ;i <1000;i++ {
		wg.Add(1)

		go func() {
			mx.Lock()			
			n++
			mx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	
	return int(n)

}

// Similary there is sometign called RWMutex where you have ReadLock too
/**
  mx.RLock() -> Lock for reading
  mx.RUnlock() -> for relieving the lock. 

  mx.Lock() -> for normal wirting
  mx.Unlock() -> For normal unlocking. 

*/


func doWithSyncAtomic() int {

	var n int64
	
	var wg sync.WaitGroup


	for i:=0 ;i <1000;i++ {
		wg.Add(1)

		go func() {
			atomic.AddInt64(&n, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	
	return int(n)

}