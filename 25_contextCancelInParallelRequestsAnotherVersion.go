package main

import (
	
	"net/http"
	"time"
	"log"
	"context"
	"runtime"

)

type Result struct {
	url string
	err error
    latency time.Duration
}

func getData(ctx context.Context, url string , ch chan<- Result) {
	startTime := time.Now()
	var r Result

	ticker := time.NewTicker(1 * time.Second).C
	 req,_ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req) ; err != nil {
		r = Result{ url, err, 0}
	} else {
		duration := time.Since(startTime).Round(time.Millisecond)
		r = Result{url, nil, duration}
		resp.Body.Close()
	}

	for {
		select {
		case  ch<-r:
			return
		case <-ticker:
			log.Println(" TICK -> ", r)
		}
	}

}

/**
This function is the one that will return the first data it is going to recieve out of all the URL's. 

*/
func first(ctx context.Context, urls []string) (*Result, error){

	//results := make(chan Result) This will call the other goRoutines to block 
	results := make(chan Result, len(urls))
	ctx , cancel := context.WithCancel(ctx)

	defer cancel()

	for _, url := range urls {
		go getData(ctx, url, results)
	}

	select {
	case res :=<- results:
		return &res, nil

	// This case can be used tby any context above this subtree of context . What if he cancels it
	case <-ctx.Done():  
		return nil,ctx.Err()
	}

}


func main(){

	//ch := make(chan Result)
	urls := [] string {
		"https://yahoo.com",
		"https://amarujala.com",
		"https://jagran.com",
		"https://timesofindia.com",
		"https://httpbin.org/delay/10", // This URL will delay response for 10 seconds. 

	}


	r, _ := first(context.Background(), urls)

	log.Println(r)
	time.Sleep(10 * time.Second)
	log.Println(" Numb of Routines Runing ", runtime.NumGoroutine())
	/*
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second) // 5 seconds timeout

	defer cancel()

	for _,url := range urls {
		go getData(ctx, url, ch)

	}

	for range urls {

		rs := <-ch

		if rs.err != nil {
			log.Printf("URL => %30s Time %s   Err!!! %s", rs.url, rs.latency, rs.err)	

		} else {
			log.Printf("URL => %-30s Time %s ", rs.url, rs.latency)	
		}


	}*/


}