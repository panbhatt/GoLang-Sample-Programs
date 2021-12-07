package main

import (
	
	"net/http"
	"time"
	"log"
	"context"

)

type Result struct {
	url string
	err error
    latency time.Duration
}

func getData(ctx context.Context, url string , ch chan<- Result) {
	startTime := time.Now()

	 req,_ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req) ; err != nil {
		ch <- Result{ url, err, 0}
	} else {
		duration := time.Since(startTime).Round(time.Millisecond)
		ch<-Result{url, nil, duration}
		resp.Body.Close()
	}

}


func main(){

	ch := make(chan Result)
	urls := [] string {
		"https://yahoo.com",
		"https://amarujala.com",
		
		"https://jagran.com",

		"https://timesofindia.com",
		"https://httpbin.org/delay/10", // This URL will delay response for 10 seconds. 

	}

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


	}


}